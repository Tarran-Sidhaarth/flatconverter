// Package runner defines the high-level interface for executing tasks based on a config file.
// It abstracts the core logic of the Buffman application, orchestrating the
// conversion and generation workflows for multiple buffer types and languages.
package runner

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/machanirobotics/buffman/internal/configuration"
	"github.com/machanirobotics/buffman/internal/generate"
	"github.com/machanirobotics/buffman/internal/parser"
	"github.com/machanirobotics/buffman/internal/remote"
	"github.com/machanirobotics/buffman/internal/utilities"
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

	parseOptions, err := r.getParserOptions(config)
	if err != nil {
		return err
	}
	if err := r.Parser.ConvertAll(ctx, parseOptions); err != nil {
		return err
	}

	generateOptions, err := r.getGenerateOptions(config)
	if err != nil {
		return err
	}

	if err := r.Generate.GenerateAll(ctx, generateOptions); err != nil {
		return err
	}

	// add other runners later
	return nil
}

func (r *runnerImpl) initializeRunner(config *configuration.Config) error {
	rem, err := remote.NewRemote(remote.Github)
	if err != nil {
		return err
	}
	source := r.getSource(config)

	for _, input := range config.Inputs {
		if input.Remote != "" {

			googleRepo := ""
			if strings.Contains(input.Name, "google") {
				googleRepo = "google"
			}

			remotePath := path.Join(source.Path, googleRepo, strings.Split(input.Remote, "/")[len(strings.Split(input.Remote, "/"))-1])
			if input.Name == source.Name {
				remotePath = input.Path
			}

			if err := rem.Pull(remote.PullOptions{Out: remotePath, Url: input.Remote, Commit: &input.Commit}); err != nil {
				return err
			}

			if strings.Contains(input.Remote, "https://github.com/protocolbuffers/protobuf") {
				if err := utilities.HandleGoogleProtobufFiles(remotePath); err != nil {
					return err
				}
			}
		}
	}

	parserManager := parser.NewManager()
	if err := parserManager.RegisterParsers(parser.Flatbuffers); err != nil { // add other parsers once it is there, this is statically defined
		return err
	}

	generateManager := generate.NewManager()
	if err := generateManager.RegisterGenerate(generate.Flatbuffers); err != nil {
		return err
	}

	r.Parser = parserManager
	r.Generate = generateManager
	r.ProtoDir = source.Path

	return nil
}
