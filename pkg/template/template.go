package template

import (
	"fmt"
	"strings"
)

type Template struct {
	Base string
}

// Format formats the template using fmt.Sprintf with dynamic fields
func (t *Template) Format(fields ...any) string {
	return fmt.Sprintf(t.Base, fields...)
}

// FormatWithMap formats the template using named placeholders
func (t *Template) FormatWithMap(data map[string]any) string {
	result := t.Base
	for key, value := range data {
		placeholder := fmt.Sprintf("{%s}", key)
		result = strings.ReplaceAll(result, placeholder, fmt.Sprintf("%v", value))
	}
	return result
}

// NewTemplate creates a new Template instance
func NewTemplate(base string) *Template {
	return &Template{
		Base: base,
	}
}
