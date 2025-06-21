package runner

import (
	"context"
	"fmt"

	"github.com/machanirobotics/buffman/pkg/converter"
	"github.com/machanirobotics/buffman/pkg/generator"
)

type Runner interface {
	Run(ctx context.Context, filePath string) error
}

func NewRunner() Runner {
	return &runnerImpl{}
}

type runnerImpl struct{}

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

func (r *runnerImpl) runConversion(ctx context.Context, convertConfig *ConvertConfig) error {
	if convertConfig == nil {
		return nil
	}

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

func (r *runnerImpl) runGeneration(ctx context.Context, config *Config) error {
	if config.Generate == nil {
		return nil
	}

	generateConfig := config.Generate.Flatbuffers

	return r.generateForLanguages(ctx, &generateConfig)
}

func (r *runnerImpl) generateForLanguages(ctx context.Context, generateConfig *FlatbuffersGenerate) error {
	languageMap := map[*LanguageConfig]generator.Languages{
		generateConfig.Languages.Cpp:    generator.CPP,
		generateConfig.Languages.Go:     generator.GO,
		generateConfig.Languages.Java:   generator.JAVA,
		generateConfig.Languages.Kotlin: generator.KOTLIN,
	}

	for langConfig, lang := range languageMap {
		if langConfig != nil {
			// Create a new generator instance for each language with its specific output directory
			gen, err := generator.NewGenerator(
				generator.FLATBUFFER,
				generateConfig.InputDir,
				langConfig.OutputDir, // Use language-specific output directory
				map[generator.Languages]string{
					lang: langConfig.ModuleOptions,
				},
			)
			if err != nil {
				return fmt.Errorf("failed to create generator for %v: %w", lang, err)
			}

			if err := gen.Generate(ctx, []generator.Languages{lang}); err != nil {
				return fmt.Errorf("failed to generate for %v: %w", lang, err)
			}
		}
	}

	return nil
}
