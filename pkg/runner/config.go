// Package runner defines the configuration structures and logic for the Buffman CLI.
// It uses the Koanf library to load, merge, and validate settings from a YAML
// configuration file, and then orchestrates the conversion and generation tasks.
// The configuration supports multiple buffer types (FlatBuffers, NanoBuffers, etc.)
// with multiple target languages, each having their own output directories and options.
package runner

import (
	"fmt"
	"path/filepath"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Config is the top-level structure that maps directly to the buffman.yml file.
// It defines the version, input directories, and plugins for code generation.
type Config struct {
	Version string   `koanf:"version"` // Configuration file format version
	Input   Input    `koanf:"input"`   // Source directories for schema files
	Plugins []Plugin `koanf:"plugins"` // Code generation plugins configuration
}

// Input defines the source directories for schema files.
// Each input specifies a directory containing schema files to be processed.
type Input struct {
	Directory string `koanf:"directory"` // Path to the directory containing schema files
}

// Plugin represents a code generation plugin (e.g., flatbuffers, nanobuffers).
// Each plugin can generate code for multiple target languages with their own configurations.
type Plugin struct {
	Name      string             `koanf:"name"`      // Plugin name (e.g., "flatbuffers", "nanobuffers")
	Out       string             `koanf:"out"`       // Output directory of the schema
	Languages []LanguageGenerate `koanf:"languages"` // Target languages and their configurations
}

// LanguageGenerate defines the configuration for generating code in a specific language.
// It specifies the target language, output directory, and language-specific options.
type LanguageGenerate struct {
	Language string `koanf:"language"` // Target language name (e.g., "go", "cpp", "java")
	Out      string `koanf:"out"`      // Output directory for generated files
	Opt      string `koanf:"opt"`      // Language-specific options (e.g., package prefix, module name)
}

// GetPluginByName searches for a plugin with the specified name in the configuration.
// It returns a pointer to the plugin if found, or nil if no plugin with the given name exists.
func (c *Config) GetPluginByName(name string) *Plugin {
	for i, plugin := range c.Plugins {
		if plugin.Name == name {
			return &c.Plugins[i]
		}
	}
	return nil
}

// GetFlatbuffersPlugin returns the flatbuffers plugin configuration if it exists.
// This is a convenience method for accessing the most commonly used plugin.
// Returns nil if no flatbuffers plugin is configured.
func (c *Config) GetFlatbuffersPlugin() *Plugin {
	return c.GetPluginByName("flatbuffers")
}

// GetNanobuffersPlugin returns the nanobuffers plugin configuration if it exists.
// This is a convenience method for accessing the nanobuffers plugin.
// Returns nil if no nanobuffers plugin is configured.
func (c *Config) GetNanobuffersPlugin() *Plugin {
	return c.GetPluginByName("nanobuffers")
}

// HasAnyLanguage checks if the plugin has any languages configured.
// Returns true if at least one language is configured for this plugin, false otherwise.
func (p *Plugin) HasAnyLanguage() bool {
	return len(p.Languages) > 0
}

// GetLanguageConfig returns the configuration for a specific language within the plugin.
// It searches through the plugin's configured languages and returns the matching configuration.
// Returns nil if the specified language is not configured for this plugin.
func (p *Plugin) GetLanguageConfig(language string) *LanguageGenerate {
	for i, lang := range p.Languages {
		if lang.Language == language {
			return &p.Languages[i]
		}
	}
	return nil
}

// GetConfiguredLanguages returns a map of language names to their configurations.
// The map keys are language names (e.g., "go", "cpp") and values are their configurations.
// This method is useful for iterating over all configured languages for a plugin.
func (p *Plugin) GetConfiguredLanguages() map[string]*LanguageGenerate {
	languages := make(map[string]*LanguageGenerate)
	for i, lang := range p.Languages {
		languages[lang.Language] = &p.Languages[i]
	}
	return languages
}

// LoadConfig reads the specified YAML configuration file, applies defaults, validates
// the settings, and unmarshals them into a Config struct.
//
// The process follows these steps:
// 1. Resolve the absolute path of the configuration file.
// 2. Load the hard-coded default values.
// 3. Load the user-provided YAML file, which overrides the defaults.
// 4. Unmarshal the final configuration into the Config struct.
// 5. Validate the struct to ensure all required fields are present.
//
// It returns a populated Config struct or an error if any step fails.
func LoadConfig(filename string) (*Config, error) {
	k := koanf.New(".")

	absPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve file path: %v", err)
	}

	if err := loadDefaults(k); err != nil {
		return nil, fmt.Errorf("failed to load default configuration: %v", err)
	}

	if err := k.Load(file.Provider(absPath), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("failed to load config file '%s': %v", filename, err)
	}

	var config Config
	if err := k.Unmarshal("", &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("invalid configuration: %v", err)
	}

	return &config, nil
}

// loadDefaults sets up the default configuration values using a confmap provider.
// These defaults ensure that the application has a sane starting point and users
// only need to specify values they wish to override.
// The defaults include a basic input directory and an empty plugins array.
func loadDefaults(k *koanf.Koanf) error {
	defaults := map[string]any{
		"version": "v1",
		"inputs": []map[string]any{
			{"directory": "protobuf"},
		},
		"plugins": []map[string]any{},
	}

	return k.Load(confmap.Provider(defaults, "."), nil)
}

// validateConfig serves as the main validation entry point, calling specific
// validation functions for each major section of the configuration.
// It ensures that all required fields are present and have valid values.
func validateConfig(c *Config) error {
	if c.Version == "" {
		return fmt.Errorf("'version' is required")
	}

	if c.Input.Directory == "" {
		return fmt.Errorf("input directory is required")
	}

	if len(c.Plugins) == 0 {
		return fmt.Errorf("at least one plugin must be configured")
	}

	for i, plugin := range c.Plugins {
		if err := validatePlugin(&plugin, i); err != nil {
			return err
		}
	}

	return nil
}

// validatePlugin checks that a plugin configuration is valid.
// It ensures that the plugin has a name and at least one language configured,
// and that each language configuration is complete and valid.
// The index parameter is used for generating descriptive error messages.
func validatePlugin(plugin *Plugin, index int) error {
	if plugin.Name == "" {
		return fmt.Errorf("plugins[%d].name is required", index)
	}

	if plugin.Out == "" {
		return fmt.Errorf("plugin[%d] output path is required", index)
	}

	if !plugin.HasAnyLanguage() {
		return fmt.Errorf("plugins[%d] (%s) must have at least one language configured", index, plugin.Name)
	}

	for j, lang := range plugin.Languages {
		if lang.Language == "" {
			return fmt.Errorf("plugins[%d].languages[%d].language is required", index, j)
		}
		if lang.Out == "" {
			return fmt.Errorf("plugins[%d].languages[%d].out is required for language '%s'", index, j, lang.Language)
		}
	}

	return nil
}
