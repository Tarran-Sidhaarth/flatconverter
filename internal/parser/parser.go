// Package converter provides interfaces and constructors for converting proto files
// into different target formats (e.g., FlatBuffers).
package parser

import (
	"context"
	"errors"

	"github.com/machanirobotics/buffman/internal/options"
	"github.com/machanirobotics/buffman/internal/parser/flatbuffers"
)

// ConverterType represents the type of converter to use.
type ParserType string

const (
	// CONVERTER_TYPE_FLATBUFFER specifies the FlatBuffer converter.
	Flatbuffers = "flatbuffers"
)

// Converter is the main interface for proto file conversion.
//
// Implementations of Converter should provide logic to convert proto files
// from a source directory to a target format, optionally keeping cleaned files.
type Parser interface {
	// Convert converts proto files according to the implementation.
	//
	// ctx: Context for cancellation and deadlines.
	// keepCleaned: If true, retains intermediate cleaned files.
	// Returns an error if the conversion fails.
	Parse(ctx context.Context, opts options.ParseOptions) error
}

// NewConverter returns a Converter implementation based on the provided ConverterType.
//
// converterType: The desired type of converter (e.g., CONVERTER_TYPE_FLATBUFFER).
// protoDir: Directory containing the source proto files.
// cleanedDir: Directory for storing cleaned proto files.
// targetDir: Directory for the converted output.
//
// Returns a Converter and nil error on success.
// Returns a non-nil error if the converterType is invalid.
func NewParser(parserType ParserType) (Parser, error) {
	switch parserType {
	case Flatbuffers:
		return flatbuffers.NewFlatbuffersParser()
	default:
		return nil, errors.New("unsupported parser type")
	}
}

type Manager struct {
	parsers map[ParserType]Parser
}

func NewManager() *Manager {
	return &Manager{
		parsers: make(map[ParserType]Parser),
	}
}

func (m *Manager) RegisterParsers(parserTypes ...ParserType) error {
	for _, p := range parserTypes {
		parser, err := NewParser(p)
		if err != nil {
			return err
		}
		m.parsers[p] = parser
	}
	return nil
}

func (m *Manager) GetParser(parserType ParserType) Parser {
	return m.parsers[parserType]
}

func (m *Manager) ConvertAll(ctx context.Context, parserOpts map[ParserType]options.ParseOptions) error {
	for parserType := range parserOpts {
		parser, exitst := m.parsers[parserType]
		if !exitst {
			return errors.New("unsupported parser")
		}
		if err := parser.Parse(ctx, parserOpts[parserType]); err != nil {
			return err
		}
	}
	return nil
}
