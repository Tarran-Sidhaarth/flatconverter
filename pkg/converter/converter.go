// pkg/converter/interface.go
package converter

import (
	"context"
	"fmt"
)

type ConverterType int

const (
	CONVERTER_TYPE_FLATBUFFER = iota
)

// Converter defines the main interface for proto file conversion
type Converter interface {
	Convert(ctx context.Context, keepCleaned bool) error
}

func NewConverter(converterType ConverterType, protoDir, cleanedDir, targetDir string) (Converter, error) {
	switch converterType {
	case CONVERTER_TYPE_FLATBUFFER:
		converter := NewFlatConverter(protoDir, cleanedDir, targetDir)
		return converter, nil
	default:
		return nil, fmt.Errorf("invalid converter type")
	}
}
