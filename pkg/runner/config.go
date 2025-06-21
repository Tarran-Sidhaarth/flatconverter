// package runner defines the configuration structures and logic for the Buffman CLI.
// It uses the Koanf library to load, merge, and validate settings from a YAML
// configuration file, and then orchestrates the conversion and generation tasks.
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
// It holds the configuration for all major tasks like 'convert' and 'generate'.
type Config struct {
	Convert  *ConvertConfig  `koanf:"convert"`
	Generate *GenerateConfig `koanf:"generate"`
}

// ConvertConfig holds all configuration related to the 'convert' tasks.
// It specifies the source directory for proto files and the target conversion formats.
type ConvertConfig struct {
	ProtoDir    string            `koanf:"proto_dir"`
	Flatbuffers FlatbuffersOutput `koanf:"flatbuffers"`
}

// FlatbuffersOutput defines the output settings for a FlatBuffers conversion task.
type FlatbuffersOutput struct {
	OutputDir string `koanf:"output_dir"`
}

// GenerateConfig holds all configuration related to the 'generate' tasks.
type GenerateConfig struct {
	Flatbuffers FlatbuffersGenerate `koanf:"flatbuffers"`
}

// FlatbuffersGenerate defines the settings for a FlatBuffers code generation task.
type FlatbuffersGenerate struct {
	InputDir  string    `koanf:"input_dir"`
	Languages Languages `koanf:"languages"`
}

// Languages defines the target languages for code generation. Each field corresponds
// to a language and holds its specific configuration.
type Languages struct {
	Cpp    *LanguageConfig `koanf:"cpp"`
	Go     *LanguageConfig `koanf:"go"`
	Java   *LanguageConfig `koanf:"java"`
	Kotlin *LanguageConfig `koanf:"kotlin"`
}

// LanguageConfig holds the configuration for a single target language, including
// its dedicated output directory and any language-specific module options.
type LanguageConfig struct {
	OutputDir     string `koanf:"output_dir"`
	ModuleOptions string `koanf:"module_options"`
}

// HasAnyLanguage is a helper method that checks if at least one language is
// configured for generation under the 'generate.flatbuffers.languages' key.
func (l *Languages) HasAnyLanguage() bool {
	return l.Cpp != nil || l.Go != nil || l.Java != nil || l.Kotlin != nil
}

// GetConfiguredLanguages is a helper method that returns a map of language names
// to their configurations for all languages that are explicitly defined in the config.
func (l *Languages) GetConfiguredLanguages() map[string]*LanguageConfig {
	languages := make(map[string]*LanguageConfig)

	if l.Cpp != nil {
		languages["cpp"] = l.Cpp
	}
	if l.Go != nil {
		languages["go"] = l.Go
	}
	if l.Java != nil {
		languages["java"] = l.Java
	}
	if l.Kotlin != nil {
		languages["kotlin"] = l.Kotlin
	}

	return languages
}

// LoadConfig reads the specified YAML configuration file, applies defaults, validates
// the settings, and unmarshals them into a Config struct.
//
// The process is as follows:
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
		return nil, fmt.Errorf("failed to resolve file path: %w", err)
	}

	if err := loadDefaults(k); err != nil {
		return nil, fmt.Errorf("failed to load default configuration: %w", err)
	}

	if err := k.Load(file.Provider(absPath), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("failed to load config file '%s': %w", filename, err)
	}

	var config Config
	if err := k.Unmarshal("", &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &config, nil
}

// loadDefaults sets up the default configuration values using a confmap provider.
// These defaults ensure that the application has a sane starting point and users
// only need to specify values they wish to override.
func loadDefaults(k *koanf.Koanf) error {
	defaults := map[string]any{
		"convert.proto_dir":                                    "",
		"convert.flatbuffers.output_dir":                       "./",
		"generate.flatbuffers.input_dir":                       "",
		"generate.flatbuffers.languages.cpp.output_dir":        "./generated/cpp",
		"generate.flatbuffers.languages.cpp.module_options":    "",
		"generate.flatbuffers.languages.go.output_dir":         "./generated/go",
		"generate.flatbuffers.languages.go.module_options":     "",
		"generate.flatbuffers.languages.java.output_dir":       "./generated/java",
		"generate.flatbuffers.languages.java.module_options":   "",
		"generate.flatbuffers.languages.kotlin.output_dir":     "./generated/kotlin",
		"generate.flatbuffers.languages.kotlin.module_options": "",
	}

	return k.Load(confmap.Provider(defaults, "."), nil)
}

// validateConfig serves as the main validation entry point, calling specific
// validation functions for each major section of the configuration.
func validateConfig(c *Config) error {
	if err := validateConvertConfig(c.Convert); err != nil {
		return err
	}

	if err := validateGenerateConfig(c.Generate); err != nil {
		return err
	}

	return nil
}

// validateConvertConfig checks that the [convert] section of the config is valid.
// It ensures that if the section is defined, `proto_dir` is not empty.
func validateConvertConfig(convert *ConvertConfig) error {
	if convert == nil {
		return nil // 'convert' section is optional.
	}

	if convert.ProtoDir == "" {
		return fmt.Errorf("'convert.proto_dir' is required but was not provided")
	}

	return nil
}

// validateGenerateConfig checks that the [generate] section of the config is valid.
// It performs multiple checks for the flatbuffers generation task.
func validateGenerateConfig(generate *GenerateConfig) error {
	if generate == nil {
		return nil // 'generate' section is optional.
	}

	if generate.Flatbuffers.InputDir == "" {
		return fmt.Errorf("'generate.flatbuffers.input_dir' is required but was not provided")
	}

	if !generate.Flatbuffers.Languages.HasAnyLanguage() {
		return fmt.Errorf("at least one language must be configured under 'generate.flatbuffers.languages'")
	}

	// Validate that each configured language has an output directory.
	for langName, langConfig := range generate.Flatbuffers.Languages.GetConfiguredLanguages() {
		if langConfig.OutputDir == "" {
			return fmt.Errorf("'output_dir' is missing for language '%s' in generate configuration", langName)
		}
	}

	return nil
}
