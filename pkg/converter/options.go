package converter

// CleanOptions configures the cleaning process
type CleanOptions struct {
	CreateNewDirectory bool
	OutputDirectory    string
	BackupOriginal     bool
	PreserveComments   bool
	InputPaths         []string
	ImportPaths        []string
}

// ConvertOptions configures the conversion to FlatBuffers
type ConvertOptions struct {
	OutputFormat     OutputFormat
	OutputDirectory  string
	IncludeServices  bool
	GenerateBindings []Language
}

type OutputFormat string

const (
	OutputFormatFlatBuffer OutputFormat = "flatbuffer"
	OutputFormatJSON       OutputFormat = "json"
	OutputFormatYAML       OutputFormat = "yaml"
)

type Language string

const (
	LanguageGo     Language = "go"
	LanguageJava   Language = "java"
	LanguagePython Language = "python"
)

type MessageInfo struct {
	Name   string
	Fields []FieldInfo
}

type FieldInfo struct {
	Name     string
	Type     string
	Number   int
	Repeated bool
}

type ServiceInfo struct {
	Name    string
	Methods []MethodInfo
}

type MethodInfo struct {
	Name       string
	InputType  string
	OutputType string
}
