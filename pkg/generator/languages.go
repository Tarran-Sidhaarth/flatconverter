package generator

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
)

// Languages represents the supported target programming languages for code generation[1].
type Languages int

// Enumeration of supported languages for code generation.
const (
	CPP     Languages = iota // Represents the C++ language.
	GO                       // Represents the Go language.
	JAVA                     // Represents the Java language.
	KOTLIN                   // Represents the Kotlin language.
	UNKNOWN                  // Represents an unsupported or unknown language.
)

// StringToLanguage converts a string representation of a language (e.g., "go")
// into the corresponding Languages enum constant. It is case-sensitive.
// It returns an error if the provided language string is not supported[1].
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
		return UNKNOWN, fmt.Errorf("unsupported language: %q", language)
	}
}

// generateCppFilesFromFbs is a helper method that generates C++ source files
// from the .fbs schemas by invoking the flatc compiler with C++ specific flags[2].
// The packagePrefix parameter is ignored as it is not used by the C++ generator.
func (f *FlatGenerator) generateCppFilesFromFbs(_ string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--cpp", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, "C++", ".h")
}

// generateGoFilesFromFbs is a helper method that generates Go source files
// from the .fbs schemas by invoking the flatc compiler with Go specific flags[2].
// It uses the packagePrefix to set the Go module name for the generated files.
func (f *FlatGenerator) generateGoFilesFromFbs(packagePrefix string) error {
	goOpts := ""
	if packagePrefix != "" {
		goOpts = fmt.Sprintf("--go-module-name %s", packagePrefix)
	}
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--go", f.targetDir, goOpts, path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, "Go", ".go")
}

// generateJavaFilesFromFbs is a helper method that generates Java source files
// from the .fbs schemas by invoking the flatc compiler with Java specific flags[2].
// It uses the packagePrefix to set the Java package name for the generated files.
func (f *FlatGenerator) generateJavaFilesFromFbs(packagePrefix string) error {
	javaOpts := ""
	if packagePrefix != "" {
		javaOpts = fmt.Sprintf("--java-package-prefix %s", packagePrefix)
	}
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--java", f.targetDir, javaOpts, path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, "Java", ".java")
}

// generateKotlinFilesFromFbs is a helper method that generates Kotlin source files
// from the .fbs schemas by invoking the flatc compiler with Kotlin specific flags[2].
// The packagePrefix parameter is currently unused for Kotlin generation.
func (f *FlatGenerator) generateKotlinFilesFromFbs(packagePrefix string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--kotlin", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, "Kotlin", ".kt")
}

// getIncludePaths walks the entire directory tree starting from the generator's
// base flatbuffer directory. It collects all subdirectory paths and formats them
// as `-I <path>` strings. This allows the flatc compiler to resolve imports
// between .fbs files located in different subdirectories.
// It returns a slice of formatted include path strings or an error if the
// directory walk fails.
func (f *FlatGenerator) getIncludePaths() ([]string, error) {
	var paths []string
	err := filepath.WalkDir(f.flatbufferDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Add every directory encountered to the include paths.
		if d.IsDir() {
			paths = append(paths, "-I "+path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk include directories: %w", err)
	}
	return paths, nil
}

func (f *FlatGenerator) createLanguageFiles(cmdStr string, language, extension string) error {
	if err := executeCommand(cmdStr); err != nil {
		return err
	}
	langFiles, err := listLanguageFiles(f.targetDir, extension)
	if err != nil {
		return err
	}
	for _, file := range langFiles {
		if err := f.insertGeneratedComments(language, path.Join(f.targetDir, file)); err != nil {
			return err
		}
	}
	return nil
}
