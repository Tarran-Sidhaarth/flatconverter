package configuration

import (
	"github.com/machanirobotics/buffman/internal"
	"github.com/machanirobotics/buffman/internal/generate/language"
)

type Subcommand string

const (
	Pkg Subcommand = "package"
)

type ToolOption struct {
	Flag        string
	Description string
}

type CommandOptions struct {
	Version string
	Options map[Subcommand]ToolOption
}

var CommandOptionsMap = map[language.Language]map[string]CommandOptions{
	language.Go: {
		"buffman": {
			Version: internal.Version,
			Options: map[Subcommand]ToolOption{
				Pkg: {
					Description: `this is used to specify the name of the go package
					example: go_package=github.com/machanirobotics/buffman/examples/go/fb`,
				},
			},
		},
		"flatbuffer": {
			Version: "2.0.0",
			Options: map[Subcommand]ToolOption{
				Pkg: {
					Flag: "--go-module-name",
					Description: `this is used to specify the name of the go package when flatbuffers are generated
					example: --go-module-name=github.com/machanirobotics/buffman/examples/go/fb`,
				},
			},
		},
		// add more types like nanobuffers later
	},
	// Add more languages if needed
	language.Java: {
		"flatbuffer": {
			Version: "2.0.0",
			Options: map[Subcommand]ToolOption{
				"package": {
					Flag: "--java-package-prefix",
					Description: `this is used to specify the prefix of the java package when flatbuffers are generated
					example: --go-module-name=com.machanirobotics.fb`,
				},
			},
		},
	},
}
