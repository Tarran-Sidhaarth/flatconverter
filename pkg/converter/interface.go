// pkg/converter/interface.go
package converter

import "context"

// Converter defines the main interface for proto file conversion
type converter interface {
	Clean(ctx context.Context) error
	Convert(ctx context.Context) error
}
