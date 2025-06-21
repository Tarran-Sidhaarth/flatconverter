// package generator provides interfaces and implementations for generating code
// from various schema definition files. It uses a factory pattern to create
// specific generator instances.
package generator

import (
	"context"
	"fmt"
)

// GeneratorType is an enumeration for different types of code generators.
// It is used in the factory function to specify which generator to create.
type GeneratorType int

const (
	// FLATBUFFER represents the FlatBuffers code generator.
	FLATBUFFER = iota
	// Add other generator types here in the future.
)

// Generator defines the common interface for all code generators. Each generator
// implementation must provide a method to perform the generation task.
type Generator interface {
	// Generate executes the code generation process for a given slice of target languages.
	// The context can be used to handle cancellation or timeouts of the generation process.
	Generate(ctx context.Context, languages []Languages) error
}

// NewGenerator acts as a factory function that constructs and returns a specific
// implementation of the Generator interface based on the provided GeneratorType.
// This is the primary entry point for acquiring a configured generator instance.
//
// Parameters:
//   - generatorType: The type of generator to create (e.g., FLATBUFFER).
//   - flatbufferDir: The source directory containing schema files (e.g., .fbs files).
//   - targetDir: The base output directory where generated code will be placed.
//   - packagePrefix: A map of language-specific options, such as package names or module prefixes.
//
// It returns a configured Generator instance or an error if the generatorType is
// unsupported or if the initialization of the specific generator fails.
func NewGenerator(generatorType GeneratorType, flatbufferDir, targetDir string, packagePrefix map[Languages]string) (Generator, error) {
	switch generatorType {
	case FLATBUFFER:
		// newFlatGenerator is defined in another file, e.g., flat_generator.go
		return newFlatGenerator(flatbufferDir, targetDir, packagePrefix)
	default:
		return nil, fmt.Errorf("invalid generator type specified: %d", generatorType)
	}
}
