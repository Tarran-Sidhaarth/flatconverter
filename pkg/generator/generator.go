package generator

import (
	"context"
	"fmt"
)

type GeneratorType int

const (
	FLATBUFFER = iota
)

type Generator interface {
	Generate(ctx context.Context, languages []Languages) error
}

func NewGenerator(generatorType GeneratorType, flatbufferDir, targetDir string, packagePrefix map[Languages]string) (Generator, error) {
	switch generatorType {
	case FLATBUFFER:
		return newFlatGenerator(flatbufferDir, targetDir, packagePrefix)
	default:
		return nil, fmt.Errorf("invalid type")
	}
}
