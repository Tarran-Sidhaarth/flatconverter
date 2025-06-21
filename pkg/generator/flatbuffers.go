package generator

import (
	"context"
	"fmt"
	"strings"
)

// FlatGenerator manages the generation of language-specific source code from
// FlatBuffer schema (.fbs) files. It holds the configuration required to
// invoke the FlatBuffers compiler (flatc).
type FlatGenerator struct {
	// flatbufferDir specifies the input directory containing the source .fbs schema files.
	flatbufferDir string
	// targetDir specifies the base output directory for the generated code.
	// Each language's output will be placed in a subdirectory (e.g., targetDir/go).
	targetDir string
	// packagePrefix maps a language to its specific module/package options,
	// such as a Go package path or a Java package name.
	packagePrefix map[Languages]string
	// includePaths is a space-separated string of directories to be used by
	// the flatc compiler for resolving imports (`-I` flag).
	includePaths string
}

// newFlatGenerator creates and initializes a new FlatGenerator instance.
// It takes the source directory for .fbs files, a base target directory for
// generated code, and a map of language-specific package options.
// It automatically determines and populates the necessary include paths for the
// FlatBuffers compiler. It returns an error if the include paths cannot be determined.
func newFlatGenerator(flatbufferDir, targetDir string, packagePrefix map[Languages]string) (*FlatGenerator, error) {
	flat := &FlatGenerator{
		flatbufferDir: flatbufferDir,
		targetDir:     targetDir,
		packagePrefix: packagePrefix,
	}

	// Determine all necessary include paths based on the flatbuffer directory.
	includes, err := flat.getIncludePaths()
	if err != nil {
		return nil, fmt.Errorf("failed to determine include paths: %w", err)
	}
	flat.includePaths = strings.Join(includes, " ")

	return flat, nil
}

// Generate executes the code generation process for a list of target languages.
// It iterates through the provided slice and calls the corresponding language-specific
// generation method. The context parameter is included for cancellation and timeout handling.
// If any language generation fails or an unsupported language is provided, the process
// stops and returns an error.
func (f *FlatGenerator) Generate(ctx context.Context, languages []Languages) error {
	for _, language := range languages {
		// Ensure the context has not been cancelled before starting a new task.
		if err := ctx.Err(); err != nil {
			return err
		}

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
			err = fmt.Errorf("unsupported language for generation: %v", language)
		}
		if err != nil {
			return fmt.Errorf("generation failed for language: %v", err)
		}
	}
	return nil
}
