package converter

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/bufbuild/protocompile"
	"github.com/bufbuild/protocompile/reporter"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Converter struct {
	compiler       *protocompile.Compiler
	BasePath       string
	DestinationDir string
}

// NewConverter creates a new converter instance
func NewConverter(basePath, destinationDir string) *Converter {
	return &Converter{
		compiler: &protocompile.Compiler{
			Resolver: &protocompile.SourceResolver{
				ImportPaths: []string{basePath},
			},
			Reporter: reporter.NewReporter(nil, nil),
		},
		BasePath:       basePath,
		DestinationDir: destinationDir,
	}
}

func (c *Converter) getAllProtoFiles() ([]string, error) {
	var paths []string
	err := filepath.Walk(c.BasePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		// Only skip Google API files and google/protobuf/descriptor.proto
		if strings.Contains(path, "google/api") ||
			strings.Contains(path, "google/protobuf/descriptor.proto") {
			return nil
		}
		if strings.HasSuffix(path, ".proto") {
			// Convert to relative path for compilation
			relativePath := strings.TrimPrefix(path, c.BasePath)
			relativePath = strings.TrimPrefix(relativePath, string(filepath.Separator))
			paths = append(paths, relativePath)
		}
		return nil
	})
	return paths, err
}

func (c *Converter) Clean(ctx context.Context) error {
	files, err := c.getAllProtoFiles()
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("no proto files found")
	}

	fileDetails, err := c.processProtoFiles(ctx, files)
	if err != nil {
		return err
	}

	// Create files and necessary directories
	for path, content := range fileDetails {
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
		if err := os.WriteFile(path, content, 0o644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", path, err)
		}
		fmt.Printf("Created: %s\n", path)
	}
	return nil
}

func (c *Converter) processProtoFiles(ctx context.Context, files []string) (map[string][]byte, error) {
	fileDetails := make(map[string][]byte)
	fds, err := c.compiler.Compile(ctx, files...)
	if err != nil {
		return nil, err
	}

	for _, fd := range fds {
		var builder strings.Builder

		// Write syntax
		builder.WriteString(fmt.Sprintf("syntax = \"%s\";\n\n", fd.Syntax()))

		// Write package
		if fd.Package() != "" {
			builder.WriteString(fmt.Sprintf("package %s;\n\n", fd.Package()))
		}

		// Write imports excluding Google API imports and google/protobuf/descriptor.proto
		imports := fd.Imports()
		hasImports := false
		for i := 0; i < imports.Len(); i++ {
			imp := imports.Get(i)
			importPath := string(imp.Path())
			// Skip google/api imports and google/protobuf/descriptor.proto
			if !strings.HasPrefix(importPath, "google/api") &&
				!strings.Contains(importPath, "google/protobuf/descriptor.proto") {
				builder.WriteString(fmt.Sprintf("import \"%s\";\n", importPath))
				hasImports = true
			}
		}
		if hasImports {
			builder.WriteString("\n")
		}

		// Write file options exactly as they are (no modifications)
		opts := fd.Options().(*descriptorpb.FileOptions)
		if opts != nil {
			// Handle go_package option (keep original)
			if opts.GoPackage != nil {
				builder.WriteString(fmt.Sprintf("option go_package = \"%s\";\n", *opts.GoPackage))
			}

			// Handle java_multiple_files option
			if opts.JavaMultipleFiles != nil {
				builder.WriteString(fmt.Sprintf("option java_multiple_files = %t;\n", *opts.JavaMultipleFiles))
			}

			// Handle java_outer_classname option
			if opts.JavaOuterClassname != nil {
				builder.WriteString(fmt.Sprintf("option java_outer_classname = \"%s\";\n", *opts.JavaOuterClassname))
			}

			// Handle java_package option (keep original)
			if opts.JavaPackage != nil {
				builder.WriteString(fmt.Sprintf("option java_package = \"%s\";\n", *opts.JavaPackage))
			}
		}
		builder.WriteString("\n")

		// Write enums
		enums := fd.Enums()
		for i := 0; i < enums.Len(); i++ {
			enum := enums.Get(i)
			c.writeEnum(&builder, enum)
		}

		// Write messages
		messages := fd.Messages()
		for i := 0; i < messages.Len(); i++ {
			msg := messages.Get(i)
			c.writeMessage(&builder, msg, 0)
		}

		// Write services - special handling for google/longrunning
		services := fd.Services()
		for i := 0; i < services.Len(); i++ {
			svc := services.Get(i)
			c.writeService(&builder, svc, fd)
		}

		// Store the built proto content in map with key as DestinationDir + fd.Path()
		key := filepath.Join(c.DestinationDir, fd.Path())
		fileDetails[key] = []byte(builder.String())
	}

	return fileDetails, nil
}

func (c *Converter) writeEnum(builder *strings.Builder, enum protoreflect.EnumDescriptor) {
	builder.WriteString(fmt.Sprintf("enum %s {\n", enum.Name()))

	values := enum.Values()
	for i := 0; i < values.Len(); i++ {
		value := values.Get(i)
		builder.WriteString(fmt.Sprintf("  %s = %d;\n", value.Name(), value.Number()))
	}
	builder.WriteString("}\n\n")
}

func (c *Converter) writeMessage(builder *strings.Builder, msg protoreflect.MessageDescriptor, indent int) {
	indentStr := strings.Repeat("  ", indent)

	builder.WriteString(fmt.Sprintf("%smessage %s {\n", indentStr, msg.Name()))

	// Write nested enums
	nestedEnums := msg.Enums()
	for i := 0; i < nestedEnums.Len(); i++ {
		nestedEnum := nestedEnums.Get(i)
		builder.WriteString(fmt.Sprintf("%s  enum %s {\n", indentStr, nestedEnum.Name()))
		values := nestedEnum.Values()
		for j := 0; j < values.Len(); j++ {
			value := values.Get(j)
			builder.WriteString(fmt.Sprintf("%s    %s = %d;\n", indentStr, value.Name(), value.Number()))
		}
		builder.WriteString(fmt.Sprintf("%s  }\n\n", indentStr))
	}

	// Write oneofs
	oneofs := msg.Oneofs()
	for i := 0; i < oneofs.Len(); i++ {
		oneof := oneofs.Get(i)
		builder.WriteString(fmt.Sprintf("%s  oneof %s {\n", indentStr, oneof.Name()))

		// Find fields that belong to this oneof
		fields := msg.Fields()
		for j := 0; j < fields.Len(); j++ {
			field := fields.Get(j)
			if field.ContainingOneof() == oneof {
				c.writeField(builder, field, indent+2)
			}
		}
		builder.WriteString(fmt.Sprintf("%s  }\n", indentStr))
	}

	// Write regular fields (not in oneofs)
	fields := msg.Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		if field.ContainingOneof() == nil {
			c.writeField(builder, field, indent+1)
		}
	}

	// Write nested messages
	nestedMessages := msg.Messages()
	for i := 0; i < nestedMessages.Len(); i++ {
		nested := nestedMessages.Get(i)
		c.writeMessage(builder, nested, indent+1)
	}

	builder.WriteString(fmt.Sprintf("%s}\n\n", indentStr))
}

func (c *Converter) writeField(builder *strings.Builder, field protoreflect.FieldDescriptor, indent int) {
	indentStr := strings.Repeat("  ", indent)

	// Determine field type
	fieldType := c.getFieldType(field)

	// Write field definition exactly as parsed (Google API annotations are automatically excluded)
	builder.WriteString(fmt.Sprintf("%s%s %s = %d;\n",
		indentStr, fieldType, field.Name(), field.Number()))
}

func (c *Converter) getFieldType(field protoreflect.FieldDescriptor) string {
	prefix := ""
	if field.Cardinality() == protoreflect.Repeated {
		prefix = "repeated "
	}

	var baseType string
	switch field.Kind() {
	case protoreflect.StringKind:
		baseType = "string"
	case protoreflect.Int32Kind:
		baseType = "int32"
	case protoreflect.Int64Kind:
		baseType = "int64"
	case protoreflect.Uint32Kind:
		baseType = "uint32"
	case protoreflect.Uint64Kind:
		baseType = "uint64"
	case protoreflect.BoolKind:
		baseType = "bool"
	case protoreflect.FloatKind:
		baseType = "float"
	case protoreflect.DoubleKind:
		baseType = "double"
	case protoreflect.BytesKind:
		baseType = "bytes"
	case protoreflect.MessageKind:
		// Use fully qualified name for message types
		msgDesc := field.Message()
		if msgDesc.Parent() != nil {
			// For nested messages, we need to build the full path
			baseType = c.getFullMessageName(msgDesc)
		} else {
			// For top-level messages, include package if it exists
			if pkg := msgDesc.ParentFile().Package(); pkg != "" {
				baseType = string(pkg) + "." + string(msgDesc.Name())
			} else {
				baseType = string(msgDesc.Name())
			}
		}
	case protoreflect.EnumKind:
		// Similar fix for enum types
		enumDesc := field.Enum()
		if enumDesc.Parent() != nil {
			baseType = c.getFullEnumName(enumDesc)
		} else {
			if pkg := enumDesc.ParentFile().Package(); pkg != "" {
				baseType = string(pkg) + "." + string(enumDesc.Name())
			} else {
				baseType = string(enumDesc.Name())
			}
		}
	default:
		baseType = "string"
	}

	return prefix + baseType
}

// Helper function to get full message name including nested path
func (c *Converter) getFullMessageName(msgDesc protoreflect.MessageDescriptor) string {
	var parts []string

	// Build the path from the message up to the file
	current := msgDesc
	for current != nil {
		parts = append([]string{string(current.Name())}, parts...)
		parent := current.Parent()
		if parent == nil {
			break
		}
		if parentMsg, ok := parent.(protoreflect.MessageDescriptor); ok {
			current = parentMsg
		} else {
			break
		}
	}

	// Add package if it exists
	if pkg := msgDesc.ParentFile().Package(); pkg != "" {
		return string(pkg) + "." + strings.Join(parts, ".")
	}

	return strings.Join(parts, ".")
}

// Helper function to get full enum name including nested path
func (c *Converter) getFullEnumName(enumDesc protoreflect.EnumDescriptor) string {
	var parts []string

	// Add the enum name
	parts = append(parts, string(enumDesc.Name()))

	// Check if it's nested in a message
	if parent := enumDesc.Parent(); parent != nil {
		if parentMsg, ok := parent.(protoreflect.MessageDescriptor); ok {
			parentName := c.getFullMessageName(parentMsg)
			return parentName + "." + strings.Join(parts, ".")
		}
	}

	// Add package if it exists
	if pkg := enumDesc.ParentFile().Package(); pkg != "" {
		return string(pkg) + "." + strings.Join(parts, ".")
	}

	return strings.Join(parts, ".")
}

func (c *Converter) writeService(builder *strings.Builder, svc protoreflect.ServiceDescriptor, fd protoreflect.FileDescriptor) {
	// Special handling for google/longrunning - only write messages, not RPC calls
	if strings.Contains(string(fd.Package()), "google.longrunning") {
		// Skip writing the service for longrunning files
		return
	}

	builder.WriteString(fmt.Sprintf("service %s {\n", svc.Name()))

	methods := svc.Methods()
	for i := 0; i < methods.Len(); i++ {
		method := methods.Get(i)
		builder.WriteString(fmt.Sprintf("  rpc %s(%s) returns (%s);\n",
			method.Name(),
			method.Input().Name(),
			method.Output().Name()))
	}
	builder.WriteString("}\n\n")
}
