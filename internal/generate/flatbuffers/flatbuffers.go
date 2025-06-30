package flatbuffers

import (
	"context"
	"strings"

	"github.com/machanirobotics/buffman/internal/generate/language"
	"github.com/machanirobotics/buffman/internal/options"
)

type FlatbufferGenerate struct {
}

func NewFlatbuffersGenerate() *FlatbufferGenerate {
	return &FlatbufferGenerate{}
}

func (f *FlatbufferGenerate) Generate(ctx context.Context, opts options.GenerateOptions) error {
	for lang := range opts.LanguagDetails {
		langDetails := opts.LanguagDetails[lang]
		metadata, err := language.GetMetadata(lang)
		if err != nil {
			return err
		}
		if err := generateLanguageFile(opts.InputDir, langDetails.OutputDir, strings.Join(langDetails.Opts, ""), metadata); err != nil {
			return err
		}
	}
	return nil
}
