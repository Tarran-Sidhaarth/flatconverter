// Package runner defines the high-level interface for executing tasks based on a config file.
// It abstracts the core logic of the Buffman application, orchestrating the
// conversion and generation workflows for multiple buffer types and languages.
package runner

import (
	"context"
	"fmt"
	"strings"

	"github.com/machanirobotics/buffman/pkg/generator"
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
type runnerImpl struct{}

// Run loads the configuration from the specified file path and executes all
// configured plugins in sequence. It processes each input directory and
// runs all configured plugins for code generation.
// Returns a wrapped error if any task fails.
func (r *runnerImpl) Run(ctx context.Context, filePath string) error {
	config, err := LoadConfig(filePath)
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	if err := r.runPlugins(ctx, config.Input.Directory, config.Plugins); err != nil {
		return fmt.Errorf("failed to process input directory '%s': %v", config.Input.Directory, err)
	}

	return nil
}

// runPlugins executes all configured plugins for the given input directory.
// Each plugin is processed independently, allowing for different buffer types
// and language combinations to be generated from the same source.
func (r *runnerImpl) runPlugins(ctx context.Context, inputDir string, plugins []Plugin) error {
	for _, plugin := range plugins {
		// Check if context has been cancelled before processing each plugin
		if err := ctx.Err(); err != nil {
			return err
		}

		if err := r.runPlugin(ctx, inputDir, &plugin); err != nil {
			return fmt.Errorf("failed to run plugin '%s': %v", plugin.Name, err)
		}
	}

	return nil
}

// runPlugin executes a single plugin for the given input directory.
// It determines the plugin type and delegates to the appropriate handler.
// Currently supports flatbuffers and nanobuffers plugins.
func (r *runnerImpl) runPlugin(ctx context.Context, inputDir string, plugin *Plugin) error {
	switch strings.ToLower(plugin.Name) {
	case "flatbuffers":
		return r.runFlatbuffersPlugin(ctx, inputDir, plugin)
	case "nanobuffers":
		return r.runNanobuffersPlugin(ctx, inputDir, plugin)
	default:
		return fmt.Errorf("unsupported plugin type: %s", plugin.Name)
	}
}

// runNanobuffersPlugin processes the nanobuffers plugin configuration.
// This is a placeholder for future nanobuffers implementation.
// Currently returns an error indicating the feature is not yet implemented.
func (r *runnerImpl) runNanobuffersPlugin(ctx context.Context, inputDir string, plugin *Plugin) error {
	// Placeholder for nanobuffers implementation
	return fmt.Errorf("nanobuffers plugin is not yet implemented")
}

// createLanguageEnumMap creates a mapping from language name strings to generator language enums.
// This mapping is used to convert configuration language names to the appropriate
// generator constants for code generation.
func (r *runnerImpl) createLanguageEnumMap() map[string]generator.Languages {
	return map[string]generator.Languages{
		"cpp":    generator.CPP,
		"go":     generator.GO,
		"java":   generator.JAVA,
		"kotlin": generator.KOTLIN,
		"lua":    generator.LUA,
		"php":    generator.PHP,
		"swift":  generator.SWIFT,
		"dart":   generator.DART,
		"csharp": generator.CSHARP,
		"python": generator.PYTHON,
		"rust":   generator.RUST,
		"ts":     generator.TS,
		"nim":    generator.NIM,
	}
}
