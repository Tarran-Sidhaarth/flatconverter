package generator

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
)

type Languages int

const (
	CPP = iota
	GO
	JAVA
	KOTLIN
	UNKOWN
)

func StringToLanguage(language string) (Languages, error) {
	switch language {
	case "cpp":
		return CPP, nil
	case "go":
		return GO, nil
	case "java":
		return JAVA, nil
	case "kotlin":
		return KOTLIN, nil
	default:
		return UNKOWN, fmt.Errorf("unsupported language")
	}
}

func (f *FlatGenerator) generateCppFilesFromFbs(_ string) error {
	// Construct the flatc command string
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "-c", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))

	return executeCommand(cmdStr)
}

func (f *FlatGenerator) generateGoFilesFromFbs(packagePrefix string) error {
	// Construct the flatc command string
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--go", f.targetDir, fmt.Sprintf("--go-module-name %s", packagePrefix), path.Join(f.flatbufferDir, "**/*.fbs"))
	return executeCommand(cmdStr)
}

func (f *FlatGenerator) generateJavaFilesFromFbs(packagePrefix string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--java", f.targetDir, fmt.Sprintf("--java-package-prefix %s", packagePrefix), path.Join(f.flatbufferDir, "**/*.fbs"))
	return executeCommand(cmdStr)
}

func (f *FlatGenerator) generateKotlinFilesFromFbs(packagePrefix string) error {
	// Construct the flatc command string
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--kotlin", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return executeCommand(cmdStr)
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
