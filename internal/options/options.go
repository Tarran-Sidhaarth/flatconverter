package options

import "github.com/machanirobotics/buffman/internal/generate/language"

type GenerateOptions struct {
	InputDir       string
	LanguagDetails map[language.Language]LanguageGenerateOptions
}

type ParseOptions struct {
	InputDir  string
	OutputDir string
}

type LanguageGenerateOptions struct {
	OutputDir string
	Opts      []string
}
