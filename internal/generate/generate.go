// package generator provides interfaces and implementations for generating code
// from various schema definition files. It uses a factory pattern to create
// specific generator instances.
package generate

import (
	"context"
	"errors"

	"github.com/machanirobotics/buffman/internal/generate/flatbuffers"
	"github.com/machanirobotics/buffman/internal/options"
)

// GeneratorType is an enumeration for different types of code generators.
// It is used in the factory function to specify which generator to create.
type GenerateType string

const (
	// FLATBUFFER represents the FlatBuffers code generator.
	Flatbuffers GenerateType = "flatbuffers"
	// Add other generator types here in the future.
)

// Generator defines the common interface for all code generators. Each generator
// implementation must provide a method to perform the generation task.
type Generate interface {
	// Generate executes the code generation process for a given slice of target languages.
	// The context can be used to handle cancellation or timeouts of the generation process.
	Generate(ctx context.Context, opts options.GenerateOptions) error
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
func NewGenerate(generateType GenerateType) (Generate, error) {
	switch generateType {
	case Flatbuffers:
		// newFlatGenerator is defined in another file, e.g., flat_generator.go
		return flatbuffers.NewFlatbuffersGenerate(), nil
	default:
		return nil, errors.New("unsupported generate type")
	}
}

type Manager struct {
	generators map[GenerateType]Generate
}

func NewManager() *Manager {
	return &Manager{generators: make(map[GenerateType]Generate)}
}

func (m *Manager) RegisterGenerate(generateTypes ...GenerateType) error {
	for _, g := range generateTypes {
		generate, err := NewGenerate(g)
		if err != nil {
			return err
		}
		m.generators[g] = generate
	}
	return nil
}

func (m *Manager) GetGenerate(generateType GenerateType) Generate {
	return m.generators[generateType]
}

func (m *Manager) GenerateAll(ctx context.Context, generateOpts map[GenerateType]options.GenerateOptions) error {
	for generateType := range generateOpts {
		generate, exists := m.generators[generateType]
		if !exists {
			return errors.New("unsupported generate type")
		}
		if err := generate.Generate(ctx, generateOpts[generateType]); err != nil {
			return err
		}
	}
	return nil
}
