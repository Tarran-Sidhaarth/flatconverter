package generator

import (
	"context"
	"fmt"
	"strings"
)

type FlatGenerator struct {
	flatbufferDir string
	targetDir     string // the generated files will go under targetDir/{language}
	packagePrefix map[Languages]string
	includePaths  string
}

func newFlatGenerator(flatbufferDir, targetDir string, packagePrefix map[Languages]string) (*FlatGenerator, error) {
	flat := &FlatGenerator{
		flatbufferDir: flatbufferDir,
		targetDir:     targetDir,
		packagePrefix: packagePrefix,
	}
	includes, err := flat.getIncludePaths()
	if err != nil {
		return nil, err
	}
	flat.includePaths = strings.Join(includes, " ")
	return flat, nil
}

func (f *FlatGenerator) Generate(ctx context.Context, languages []Languages) error {
	for _, language := range languages {
		var err error
		switch language {
		case CPP:
			err = f.generateCppFilesFromFbs(f.packagePrefix[language])
		case GO:
			err = f.generateGoFilesFromFbs(f.packagePrefix[language])
		case JAVA:
			err = f.generateJavaFilesFromFbs(f.packagePrefix[language])
		case KOTLIN:
			err = f.generateKotlinFilesFromFbs(f.packagePrefix[language])
		default:
			err = fmt.Errorf("unsupported language")
		}
		if err != nil {
			return err
		}
	}
	return nil
}
