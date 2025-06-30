package runner

import (
	"fmt"

	"github.com/machanirobotics/buffman/internal/configuration"
	"github.com/machanirobotics/buffman/internal/generate"
	"github.com/machanirobotics/buffman/internal/generate/language"
	"github.com/machanirobotics/buffman/internal/options"
	"github.com/machanirobotics/buffman/internal/parser"
)

func (r *runnerImpl) getSource(config *configuration.Config) configuration.Input {
	for _, input := range config.Inputs {
		if input.Name == "source" {
			return input
		}
	}
	return configuration.Input{}
}

// func (r *runnerImpl) get

func (r *runnerImpl) getParserOptions(config *configuration.Config) (map[parser.ParserType]options.ParseOptions, error) {
	parserOptions := map[parser.ParserType]options.ParseOptions{}
	for _, plugin := range config.Plugins {
		parserType := parser.ParserType(plugin.Name)
		if r.Parser.GetParser(parserType) == nil {
			return nil, fmt.Errorf("unsupported pluging name %s", plugin.Name)
		}
		parserOptions[parserType] = options.ParseOptions{
			InputDir:  r.ProtoDir,
			OutputDir: plugin.Out,
		}
	}
	return parserOptions, nil
}

func (r *runnerImpl) getGenerateOptions(config *configuration.Config) (map[generate.GenerateType]options.GenerateOptions, error) {
	generateOptions := map[generate.GenerateType]options.GenerateOptions{}
	for _, plugin := range config.Plugins {
		generateType := generate.GenerateType(plugin.Name)
		if r.Generate.GetGenerate(generateType) == nil {
			return nil, fmt.Errorf("unsupported invalid plugin name %s", plugin.Name)
		}
		languageOptions, err := r.getLangnguageOptions(plugin)
		if err != nil {
			return nil, err
		}
		generateOptions[generateType] = options.GenerateOptions{
			InputDir:       plugin.Out,
			LanguagDetails: languageOptions,
		}
	}
	return generateOptions, nil
}

func (r *runnerImpl) getLangnguageOptions(plugin configuration.Plugin) (map[language.Language]options.LanguageGenerateOptions, error) {
	languageOptions := map[language.Language]options.LanguageGenerateOptions{}
	for _, lang := range plugin.Languages {
		if !language.IsSupportedLanguage(lang.Language) {
			return nil, &language.UnsupportedLanguageError{}
		}
		langType := language.Language(lang.Language)
		languageOptions[langType] = options.LanguageGenerateOptions{
			OutputDir: lang.Out,
			Opts:      lang.Opt,
		}
	}
	return languageOptions, nil
}
