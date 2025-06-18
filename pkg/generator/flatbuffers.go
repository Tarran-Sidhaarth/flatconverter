package generator

import (
	"context"
	"fmt"
	"io/fs"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const flatcCommand = "flatc %s %s -o %s %s" // include paths, language, output directory, fbs file directory

type FlatGenerator struct {
	flatbufferDir      string
	targetDir          string // the generated files will go under targetDir/{language}
	namespaceOverrides map[Languages]string
	includePaths       string
}

func NewFlatGenerator(flatbufferDir, targetDir string, namespaceOverrides map[Languages]string) (*FlatGenerator, error) {
	flat := &FlatGenerator{
		flatbufferDir:      flatbufferDir,
		targetDir:          targetDir,
		namespaceOverrides: namespaceOverrides,
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
			err = f.generateCppFilesFromFbs(f.namespaceOverrides[language])
		case GO:
			err = f.generateGoFilesFromFbs(f.namespaceOverrides[language])
		case JAVA:
			err = f.generateJavaFilesFromFbs(f.namespaceOverrides[language])
		case KOTLIN:
			err = f.generateKotlinFilesFromFbs(f.namespaceOverrides[language])
		default:
			err = fmt.Errorf("unsupported language")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *FlatGenerator) generateCppFilesFromFbs(_ string) error {
	// Prepare the output directory (targetDir + "cpp")
	outputDir := f.targetDir + "cpp"

	// Construct the flatc command string
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "-c", outputDir, path.Join(f.flatbufferDir, "**/*.fbs"))

	// Split the command string into command and arguments
	cmdArgs := strings.Fields(cmdStr)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	// Run the command and capture output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute flatc command: %v, output: %s", err, string(output))
	}
	return nil
}

func (f *FlatGenerator) generateGoFilesFromFbs(override string) error {
	return nil
}

func (f *FlatGenerator) generateJavaFilesFromFbs(override string) error {
	return nil
}

func (f *FlatGenerator) generateKotlinFilesFromFbs(override string) error {
	return nil
}

func (f *FlatGenerator) getIncludePaths() ([]string, error) {
	var paths []string
	err := filepath.WalkDir(f.flatbufferDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		paths = append(paths, "-I "+path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paths, nil
}
