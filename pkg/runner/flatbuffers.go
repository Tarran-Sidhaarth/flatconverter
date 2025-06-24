package runner

import (
	"context"
	"fmt"
	"strings"

	"github.com/machanirobotics/buffman/pkg/converter"
	"github.com/machanirobotics/buffman/pkg/generator"
)

// runFlatbuffersPlugin processes the flatbuffers plugin configuration.
// It creates generators for each configured language and executes them
// with the appropriate output directories and options.
func (r *runnerImpl) runFlatbuffersPlugin(ctx context.Context, inputDir string, plugin *Plugin) error {
	if err := r.generateFBSFiles(ctx, inputDir, plugin.Out); err != nil {
		return err
	}

	return r.generateForLanguages(ctx, plugin.Out, plugin, generator.FLATBUFFER)
}

func (r *runnerImpl) generateFBSFiles(ctx context.Context, protoDir string, flatbuffersDir string) error {
	c, err := converter.NewConverter(converter.FLATBUFFER, protoDir, flatbuffersDir, "")
	if err != nil {
		return err
	}
	if err := c.Convert(ctx); err != nil {
		return err
	}
	return nil
}

// generateForLanguages iterates over the languages defined in the plugin configuration.
// For each configured language, it creates a new generator instance with the
// appropriate language-specific output directory and module options, then runs it.
// This ensures that each language's output is handled independently.
func (r *runnerImpl) generateForLanguages(ctx context.Context, inputDir string, plugin *Plugin, generatorType generator.GeneratorType) error {
	// Create a map of language configurations to generator language enums
	languageEnumMap := r.createLanguageEnumMap()

	for _, langConfig := range plugin.Languages {
		// Check if context has been cancelled before processing each language
		if err := ctx.Err(); err != nil {
			return err
		}

		langEnum, exists := languageEnumMap[strings.ToLower(langConfig.Language)]
		if !exists {
			return fmt.Errorf("unsupported language: %s", langConfig.Language)
		}

		// Create a new generator for each language to ensure its specific
		// configuration (output dir, module options) is respected
		gen, err := generator.NewGenerator(
			generatorType,
			inputDir,
			langConfig.Out, // Use the language-specific output directory
			map[generator.Languages]string{
				langEnum: langConfig.Opt, // Use language-specific options
			},
		)
		if err != nil {
			return fmt.Errorf("failed to create generator for language %s: %w", langConfig.Language, err)
		}

		if err := gen.Generate(ctx, []generator.Languages{langEnum}); err != nil {
			return fmt.Errorf("failed to generate code for language %s: %w", langConfig.Language, err)
		}
	}

	return nil
}
