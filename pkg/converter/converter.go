// Package converter provides interfaces and constructors for converting proto files
// into different target formats (e.g., FlatBuffers).
package converter

import (
	"context"
	"fmt"
	"strings"
)

// ConverterType represents the type of converter to use.
type ConverterType int

const (
	// CONVERTER_TYPE_FLATBUFFER specifies the FlatBuffer converter.
	FLATBUFFER = iota
)

// Converter is the main interface for proto file conversion.
//
// Implementations of Converter should provide logic to convert proto files
// from a source directory to a target format, optionally keeping cleaned files.
type Converter interface {
	// Convert converts proto files according to the implementation.
	//
	// ctx: Context for cancellation and deadlines.
	// keepCleaned: If true, retains intermediate cleaned files.
	// Returns an error if the conversion fails.
	Convert(ctx context.Context) error
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
func NewConverter(converterType ConverterType, protoDir, targetDir, prefix string) (Converter, error) {
	if !strings.HasSuffix(prefix, "/") && prefix != "" {
		prefix += "/"
	}
	switch converterType {
	case FLATBUFFER:
		return newFlatConverter(protoDir, targetDir, prefix)
	default:
		return nil, fmt.Errorf("invalid converter type")
	}
}
