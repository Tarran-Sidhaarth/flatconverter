package generator

import "context"

type Languages int

const (
	CPP = iota
	GO
	JAVA
	KOTLIN
)

type Generator interface {
	Generate(ctx context.Context, languages []Languages) error
}
