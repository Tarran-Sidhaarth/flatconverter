// Package runner defines the high-level interface for executing tasks based on a config file.
// It abstracts the core logic of the Buffman application, orchestrating the
// conversion and generation workflows for multiple buffer types and languages.
package runner

import (
	"context"
	"fmt"

	"github.com/machanirobotics/buffman/internal/configuration"
	"github.com/machanirobotics/buffman/internal/generate"
	"github.com/machanirobotics/buffman/internal/generate/language"
	"github.com/machanirobotics/buffman/internal/options"
	"github.com/machanirobotics/buffman/internal/parser"
	"github.com/machanirobotics/buffman/internal/remote"
)

// Runner defines the high-level interface for executing tasks based on a config file.
// It abstracts the core logic of the Buffman application, orchestrating the
// conversion and generation workflows across multiple plugins and languages.
type Runner interface {
	// Run loads the configuration from the given file path and executes all defined
	// plugin tasks sequentially. The context can be used to cancel the entire run.
	// It returns an error if any step fails.
	Run(ctx context.Context, filePath string) error
}

// NewRunner creates a new instance of the default Runner implementation.
// The returned Runner can process configurations with multiple plugins
// and generate code for various target languages.
func NewRunner() Runner {
	return &runnerImpl{}
}

// runnerImpl is the concrete implementation of the Runner interface.
// It handles the orchestration of conversion and generation tasks
// based on the loaded configuration.
type runnerImpl struct {
	ProtoDir string
	Parser   *parser.Manager
	Generate *generate.Manager
}

// Run loads the configuration from the specified file path and executes all
// configured plugins in sequence. It processes each input directory and
// runs all configured plugins for code generation.
// Returns a wrapped error if any task fails.
func (r *runnerImpl) Run(ctx context.Context, filePath string) error {
	config, err := configuration.LoadConfig(filePath)
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}
	if err := r.initializeRunner(config); err != nil {
		return err
	}
	if err := r.Parser.ConvertAll(ctx, map[parser.ParserType]options.ParseOptions{
		parser.Flatbuffers: options.ParseOptions{
			InputDir:  r.ProtoDir,
			OutputDir: config.Plugins[0].Out,
		}}); err != nil {
		return err
	}

	languageOptions := map[language.Language]options.LanguageGenerateOptions{}
	for _, lang := range config.Plugins[0].Languages {
		languageOptions[language.Language(lang.Language)] = options.LanguageGenerateOptions{
			OutputDir: lang.Out,
			Opts:      lang.Opt,
		}
	}

	if err := r.Generate.GenerateAll(ctx, map[generate.GenerateType]options.GenerateOptions{
		generate.Flatbuffers: options.GenerateOptions{
			InputDir:       config.Plugins[0].Out,
			LanguagDetails: languageOptions,
		},
	}); err != nil {
		return err
	}

	// add other runners later
	return nil
}

func (r *runnerImpl) initializeRunner(config *configuration.Config) error {
	protoDir := ""
	rem, err := remote.NewRemote(remote.Github)
	if err != nil {
		return err
	}
	for _, input := range config.Inputs {
		if configuration.IsSource(input.Name) {
			protoDir = input.Path
		}
		if input.Remote != "" {
			var commit *string
			if input.Commit != "" {
				commit = &input.Commit
			}
			if err := rem.Pull(remote.PullOptions{Out: "./", Url: input.Remote, Commit: commit}); err != nil {
				return err
			}
		}
	}

	parserManager := parser.NewManager()
	if err := parserManager.RegisterParsers(parser.Flatbuffers); err != nil {
		return err
	}

	generateManager := generate.NewManager()
	if err := generateManager.RegisterGenerate(generate.Flatbuffers); err != nil {
		return err
	}

	r.Parser = parserManager
	r.Generate = generateManager
	r.ProtoDir = protoDir

	return nil
}
