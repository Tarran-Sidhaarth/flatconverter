package language

type Language string

const (
	Cpp     Language = "cpp"
	Go      Language = "go"
	Java    Language = "java"
	Kotlin  Language = "kotlin"
	Lua     Language = "lua"
	Php     Language = "php"
	Swift   Language = "swift"
	Dart    Language = "dart"
	Csharp  Language = "csharp"
	Python  Language = "python"
	Rust    Language = "rust"
	Ts      Language = "typescript"
	Nim     Language = "nim"
	Unknown Language = "unknown"
)

type LanguageMetadata struct {
	Language     Language
	CommentStyle string
	Extension    string
}

// LanguageMetadataMap provides metadata for all supported languages
var languageMetadataMap = map[Language]LanguageMetadata{
	Cpp: {
		Language:     Cpp,
		CommentStyle: "//",
		Extension:    ".h",
	},
	Go: {
		Language:     Go,
		CommentStyle: "//",
		Extension:    ".go",
	},
	Java: {
		Language:     Java,
		CommentStyle: "//",
		Extension:    ".java",
	},
	Kotlin: {
		Language:     Kotlin,
		CommentStyle: "//",
		Extension:    ".kt",
	},
	Lua: {
		Language:     Lua,
		CommentStyle: "--",
		Extension:    ".lua",
	},
	Php: {
		Language:     Php,
		CommentStyle: "//",
		Extension:    ".php",
	},
	Swift: {
		Language:     Swift,
		CommentStyle: "//",
		Extension:    ".swift",
	},
	Dart: {
		Language:     Dart,
		CommentStyle: "//",
		Extension:    ".dart",
	},
	Csharp: {
		Language:     Csharp,
		CommentStyle: "//",
		Extension:    ".cs",
	},
	Python: {
		Language:     Python,
		CommentStyle: "#",
		Extension:    ".py",
	},
	Rust: {
		Language:     Rust,
		CommentStyle: "//",
		Extension:    ".rs",
	},
	Ts: {
		Language:     Ts,
		CommentStyle: "//",
		Extension:    ".ts",
	},
	Nim: {
		Language:     Nim,
		CommentStyle: "#",
		Extension:    ".nim",
	},
	Unknown: {
		Language:     Unknown,
		CommentStyle: "",
		Extension:    "",
	},
}

// GetMetadata returns the metadata for a given language
func GetMetadata(lang Language) (LanguageMetadata, error) {
	metadata, exists := languageMetadataMap[lang]
	if !exists {
		return metadata, &UnsupportedLanguageError{}
	}
	return metadata, nil
}

// GetSupportedLanguages returns a slice of all supported languages
func GetSupportedLanguages() []Language {
	languages := make([]Language, 0, len(languageMetadataMap))
	for lang := range languageMetadataMap {
		if lang != Unknown {
			languages = append(languages, lang)
		}
	}
	return languages
}

func IsSupportedLanguage(language string) bool {
	_, ok := languageMetadataMap[Language(language)]
	return ok
}

// UnsupportedLanguageError represents an error for unsupported languages
type UnsupportedLanguageError struct {
}

// Error implements the error interface
func (e *UnsupportedLanguageError) Error() string {
	return "UnsupportedLanguageError: buffman does not support this language yet!"
}
