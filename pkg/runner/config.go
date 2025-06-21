package runner

import (
	"fmt"
	"path/filepath"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Convert  *ConvertConfig  `koanf:"convert"`
	Generate *GenerateConfig `koanf:"generate"`
}

type ConvertConfig struct {
	ProtoDir    string            `koanf:"proto_dir"`
	Flatbuffers FlatbuffersOutput `koanf:"flatbuffers"`
}

type FlatbuffersOutput struct {
	OutputDir string `koanf:"output_dir"`
}

type GenerateConfig struct {
	Flatbuffers FlatbuffersGenerate `koanf:"flatbuffers"`
}

type FlatbuffersGenerate struct {
	InputDir  string    `koanf:"input_dir"`
	Languages Languages `koanf:"languages"`
}

type Languages struct {
	Cpp    *LanguageConfig `koanf:"cpp"`
	Go     *LanguageConfig `koanf:"go"`
	Java   *LanguageConfig `koanf:"java"`
	Kotlin *LanguageConfig `koanf:"kotlin"`
}

// Updated language config with individual output directory
type LanguageConfig struct {
	OutputDir     string `koanf:"output_dir"`
	ModuleOptions string `koanf:"module_options"`
}

// Helper method to check if any language is configured
func (l *Languages) HasAnyLanguage() bool {
	return l.Cpp != nil || l.Go != nil || l.Java != nil || l.Kotlin != nil
}

// Helper method to get configured languages with their configs
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

func LoadConfig(filename string) (*Config, error) {
	k := koanf.New(".")

	absPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve file path: %w", err)
	}

	if err := loadDefaults(k); err != nil {
		return nil, fmt.Errorf("failed to load defaults: %w", err)
	}

	if err := k.Load(file.Provider(absPath), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	var config Config
	if err := k.Unmarshal("", &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, validateConfig(&config)
}

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

func validateConfig(c *Config) error {
	if err := validateConvertConfig(c.Convert); err != nil {
		return err
	}

	if err := validateGenerateConfig(c.Generate); err != nil {
		return err
	}

	return nil
}

func validateConvertConfig(convert *ConvertConfig) error {
	if convert == nil {
		return nil
	}

	if convert.ProtoDir == "" {
		return fmt.Errorf("proto file directory is missing")
	}

	return nil
}

func validateGenerateConfig(generate *GenerateConfig) error {
	if generate == nil {
		return nil
	}

	if generate.Flatbuffers.InputDir == "" {
		return fmt.Errorf("flatbuffer files directory is missing")
	}

	if !generate.Flatbuffers.Languages.HasAnyLanguage() {
		return fmt.Errorf("please specify at least one language")
	}

	// Validate that each configured language has an output directory
	languages := generate.Flatbuffers.Languages.GetConfiguredLanguages()
	for langName, langConfig := range languages {
		if langConfig.OutputDir == "" {
			return fmt.Errorf("output directory is missing for language: %s", langName)
		}
	}

	return nil
}
