package runner

import (
	"context"
	"fmt"

	"github.com/machanirobotics/buffman/pkg/converter"
	"github.com/machanirobotics/buffman/pkg/generator"
)

// Runner defines the high-level interface for executing tasks based on a config file.
// It abstracts the core logic of the Buffman application, orchestrating the
// conversion and generation workflows[1].
type Runner interface {
	// Run loads the configuration from the given file path and executes all defined
	// conversion and generation tasks sequentially. The context can be used to
	// cancel the entire run. It returns an error if any step fails[1].
	Run(ctx context.Context, filePath string) error
}

// NewRunner creates a new instance of the default Runner implementation.
func NewRunner() Runner {
	return &runnerImpl{}
}

// runnerImpl is the concrete implementation of the Runner interface[1].
type runnerImpl struct{}

// Run loads the configuration from the specified file path and executes the conversion
// and generation workflows in sequence. It returns a wrapped error if any task fails[1].
func (r *runnerImpl) Run(ctx context.Context, filePath string) error {
	config, err := LoadConfig(filePath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if err := r.runConversion(ctx, config.Convert); err != nil {
		return fmt.Errorf("conversion failed: %w", err)
	}

	if err := r.runGeneration(ctx, config); err != nil {
		return fmt.Errorf("generation failed: %w", err)
	}

	return nil
}

// runConversion executes the schema conversion process if a 'convert' section
// is defined in the configuration. It is a no-op if the config section is nil[1].
func (r *runnerImpl) runConversion(ctx context.Context, convertConfig *ConvertConfig) error {
	if convertConfig == nil {
		return nil
	}

	// Currently supports only proto to flatbuffer conversion.
	convert, err := converter.NewConverter(
		converter.FLATBUFFER,
		convertConfig.ProtoDir,
		convertConfig.Flatbuffers.OutputDir,
		"",
	)
	if err != nil {
		return err
	}

	return convert.Convert(ctx)
}

// runGeneration executes the code generation process if a 'generate' section
// is defined in the configuration. It is a no-op if the config section is nil[1].
func (r *runnerImpl) runGeneration(ctx context.Context, config *Config) error {
	if config.Generate == nil {
		return nil
	}

	generateConfig := config.Generate.Flatbuffers

	return r.generateForLanguages(ctx, &generateConfig)
}

// generateForLanguages iterates over the languages defined in the generation config.
// For each configured language, it creates a new generator instance with the
// appropriate language-specific output directory and module options, then runs it.
// This ensures that each language's output is handled independently[1].
func (r *runnerImpl) generateForLanguages(ctx context.Context, generateConfig *FlatbuffersGenerate) error {
	// This map facilitates iterating over the language configurations.
	languageMap := map[*LanguageConfig]generator.Languages{
		generateConfig.Languages.Cpp:    generator.CPP,
		generateConfig.Languages.Go:     generator.GO,
		generateConfig.Languages.Java:   generator.JAVA,
		generateConfig.Languages.Kotlin: generator.KOTLIN,
	}

	for langConfig, langEnum := range languageMap {
		if langConfig != nil {
			// A new generator is created for each language to ensure its specific
			// configuration (output dir, module options) is respected.
			gen, err := generator.NewGenerator(
				generator.FLATBUFFER,
				generateConfig.InputDir,
				langConfig.OutputDir, // Use the language-specific output directory.
				map[generator.Languages]string{
					langEnum: langConfig.ModuleOptions,
				},
			)
			if err != nil {
				return fmt.Errorf("failed to create generator for language %v: %w", langEnum, err)
			}

			if err := gen.Generate(ctx, []generator.Languages{langEnum}); err != nil {
				return fmt.Errorf("failed to generate code for language %v: %w", langEnum, err)
			}
		}
	}

	return nil
}
