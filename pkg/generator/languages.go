package generator

import (
	"fmt"
	"path"

	"github.com/machanirobotics/buffman/pkg/template"
)

// Languages represents the supported target programming languages for code generation.
type Languages int

// Enumeration of supported languages for code generation.
const (
	CPP     Languages = iota // Represents the C++ language.
	GO                       // Represents the Go language.
	JAVA                     // Represents the Java language.
	KOTLIN                   // Represents the Kotlin language.
	LUA                      // Represents the Lua language.
	PHP                      // Represents the PHP language.
	SWIFT                    // Represents the Swift language.
	DART                     // Represents the Dart language.
	CSHARP                   // Represents the C# language.
	PYTHON                   // Represents the Python language.
	RUST                     // Represents the Rust language.
	TS                       // Represents the TypeScript language.
	NIM                      // Represents the Nim language.
	UNKNOWN                  // Represents an unsupported or unknown language.
)

type Language struct {
	Language      Languages
	CommentPrefix string
	Extension     string
}

var languageDetails = map[Languages]*Language{
	GO:     {Language: GO, CommentPrefix: "//", Extension: ".go"},
	JAVA:   {Language: JAVA, CommentPrefix: "//", Extension: ".java"},
	CPP:    {Language: CPP, CommentPrefix: "//", Extension: ".cpp"},
	KOTLIN: {Language: KOTLIN, CommentPrefix: "//", Extension: ".kt"},
	LUA:    {Language: LUA, CommentPrefix: "--", Extension: ".lua"},
	PHP:    {Language: PHP, CommentPrefix: "//", Extension: ".php"},
	SWIFT:  {Language: SWIFT, CommentPrefix: "//", Extension: ".swift"},
	DART:   {Language: DART, CommentPrefix: "//", Extension: ".dart"},
	CSHARP: {Language: CSHARP, CommentPrefix: "//", Extension: ".cs"},
	PYTHON: {Language: PYTHON, CommentPrefix: "#", Extension: ".py"},
	RUST:   {Language: RUST, CommentPrefix: "//", Extension: ".rs"},
	TS:     {Language: TS, CommentPrefix: "//", Extension: ".ts"},
	NIM:    {Language: NIM, CommentPrefix: "#", Extension: ".nim"},
}

// FromString converts a string representation of a language (e.g., "go")
// into the corresponding Language struct. It is case-sensitive.
// Returns a Language with UNKNOWN type if the provided language string is not supported.
func newLanguageFromString(language string) *Language {
	switch language {
	case "cpp":
		return languageDetails[CPP]
	case "go":
		return languageDetails[GO]
	case "java":
		return languageDetails[JAVA]
	case "kotlin":
		return languageDetails[KOTLIN]
	case "lua":
		return languageDetails[LUA]
	case "php":
		return languageDetails[PHP]
	case "swift":
		return languageDetails[SWIFT]
	case "dart":
		return languageDetails[DART]
	case "csharp":
		return languageDetails[CSHARP]
	case "python":
		return languageDetails[PYTHON]
	case "rust":
		return languageDetails[RUST]
	case "ts":
		return languageDetails[TS]
	case "nim":
		return languageDetails[NIM]
	default:
		return &Language{Language: UNKNOWN, CommentPrefix: "", Extension: ""}
	}
}

// ToString converts the Language enum to its string representation.
// Returns the lowercase string name of the language or "unknown" for unsupported languages.
func (l *Language) ToString() string {
	switch l.Language {
	case CPP:
		return "Cpp"
	case GO:
		return "Go"
	case JAVA:
		return "Java"
	case KOTLIN:
		return "Kotlin"
	case LUA:
		return "Lua"
	case PHP:
		return "Php"
	case SWIFT:
		return "Swift"
	case DART:
		return "Dart"
	case CSHARP:
		return "Csharp"
	case PYTHON:
		return "Python"
	case RUST:
		return "Rust"
	case TS:
		return "Ts"
	case NIM:
		return "Nim"
	default:
		return "Unknown"
	}
}

func NewLanguage(language string) *Language {
	return newLanguageFromString(language)
}

// generateCppFilesFromFbs is a helper method that generates C++ source files
// from the .fbs schemas by invoking the flatc compiler with C++ specific flags.
// The packagePrefix parameter is ignored as it is not used by the C++ generator.
func (f *FlatGenerator) generateCppFilesFromFbs(_ string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--cpp", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, CPP)
}

// generateGoFilesFromFbs is a helper method that generates Go source files
// from the .fbs schemas by invoking the flatc compiler with Go specific flags.
// It uses the packagePrefix to set the Go module name for the generated files.
func (f *FlatGenerator) generateGoFilesFromFbs(packagePrefix string) error {
	goOpts := ""
	if packagePrefix != "" {
		goOpts = fmt.Sprintf("--go-module-name %s", packagePrefix)
	}
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--go", f.targetDir, goOpts, path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, GO)
}

// generateJavaFilesFromFbs is a helper method that generates Java source files
// from the .fbs schemas by invoking the flatc compiler with Java specific flags.
// It uses the packagePrefix to set the Java package name for the generated files.
func (f *FlatGenerator) generateJavaFilesFromFbs(packagePrefix string) error {
	javaOpts := ""
	if packagePrefix != "" {
		javaOpts = fmt.Sprintf("--java-package-prefix %s", packagePrefix)
	}
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--java", f.targetDir, javaOpts, path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, JAVA)
}

// generateKotlinFilesFromFbs is a helper method that generates Kotlin source files
// from the .fbs schemas by invoking the flatc compiler with Kotlin specific flags.
// The packagePrefix parameter is currently unused for Kotlin generation.
func (f *FlatGenerator) generateKotlinFilesFromFbs(packagePrefix string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--kotlin", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, KOTLIN)
}

// generateLuaFilesFromFbs is a helper method that generates Lua source files
// from the .fbs schemas by invoking the flatc compiler with Lua specific flags.
// The packagePrefix parameter is unused for Lua generation.
func (f *FlatGenerator) generateLuaFilesFromFbs(_ string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--lua", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, LUA)
}

// generatePhpFilesFromFbs is a helper method that generates PHP source files
// from the .fbs schemas by invoking the flatc compiler with PHP specific flags.
// The packagePrefix parameter is unused for PHP generation.
func (f *FlatGenerator) generatePhpFilesFromFbs(_ string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--php", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, PHP)
}

// generateSwiftFilesFromFbs is a helper method that generates Swift source files
// from the .fbs schemas by invoking the flatc compiler with Swift specific flags.
// The packagePrefix parameter is unused for Swift generation.
func (f *FlatGenerator) generateSwiftFilesFromFbs(_ string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--swift", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, SWIFT)
}

// generateDartFilesFromFbs is a helper method that generates Dart source files
// from the .fbs schemas by invoking the flatc compiler with Dart specific flags.
// The packagePrefix parameter is unused for Dart generation.
func (f *FlatGenerator) generateDartFilesFromFbs(_ string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--dart", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, DART)
}

// generateCsharpFilesFromFbs is a helper method that generates C# source files
// from the .fbs schemas by invoking the flatc compiler with C# specific flags.
// The packagePrefix parameter is unused for C# generation.
func (f *FlatGenerator) generateCsharpFilesFromFbs(_ string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--csharp", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, CSHARP)
}

// generatePythonFilesFromFbs is a helper method that generates Python source files
// from the .fbs schemas by invoking the flatc compiler with Python specific flags.
// The packagePrefix parameter is unused for Python generation.
func (f *FlatGenerator) generatePythonFilesFromFbs(_ string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--python", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, PYTHON)
}

// generateRustFilesFromFbs is a helper method that generates Rust source files
// from the .fbs schemas by invoking the flatc compiler with Rust specific flags.
// The packagePrefix parameter is unused for Rust generation.
func (f *FlatGenerator) generateRustFilesFromFbs(_ string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--rust", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, RUST)
}

// generateTsFilesFromFbs is a helper method that generates TypeScript source files
// from the .fbs schemas by invoking the flatc compiler with TypeScript specific flags.
// The packagePrefix parameter is unused for TypeScript generation.
func (f *FlatGenerator) generateTsFilesFromFbs(_ string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--ts", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, TS)
}

// generateNimFilesFromFbs is a helper method that generates Nim source files
// from the .fbs schemas by invoking the flatc compiler with Nim specific flags.
// The packagePrefix parameter is unused for Nim generation.
func (f *FlatGenerator) generateNimFilesFromFbs(_ string) error {
	cmdStr := fmt.Sprintf(flatcCommand, f.includePaths, "--nim", f.targetDir, "", path.Join(f.flatbufferDir, "**/*.fbs"))
	return f.createLanguageFiles(cmdStr, NIM)
}

func (f *FlatGenerator) createLanguageFiles(cmdStr string, language Languages) error {
	if err := executeCommand(cmdStr); err != nil {
		return err
	}
	lang := languageDetails[language]
	langFiles, err := listLanguageFiles(f.targetDir, lang.Extension)
	if err != nil {
		return err
	}
	t := template.NewTemplate(lang.CommentPrefix)
	comment := t.BuildDefaultComment(lang.ToString())
	for _, file := range langFiles {
		if err := f.insertGeneratedComments(comment, path.Join(f.targetDir, file)); err != nil {
			return err
		}
	}
	return nil
}
