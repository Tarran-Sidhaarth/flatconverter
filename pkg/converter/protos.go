// Package converter provides functions for cleaning and processing Protocol Buffer files,
// specifically removing Google API imports and reconstructing proto file content.
package converter

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/bufbuild/protocompile"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// removeGoogleAPI returns a map of cleaned proto file paths to their content,
// with Google API imports and google/protobuf/descriptor.proto removed.
//
// It compiles the given proto files, reconstructs their content (syntax, package,
// imports, options, enums, messages, services), and writes the cleaned content
// to the outputDir with the original file's relative path as the key.
//
// ctx: Context for compilation and cancellation.
// compiler: The proto compiler to use.
// files: List of proto file paths to process.
// outputDir: Directory to use as the root for output file paths.
//
// Returns a map where keys are output file paths and values are file contents.
// Returns an error if compilation fails.
func removeGoogleAPI(ctx context.Context, compiler *protocompile.Compiler, files []string, outputDir, prefix string) (map[string][]byte, error) {
	fileDetails := make(map[string][]byte)
	fds, err := compiler.Compile(ctx, files...)
	if err != nil {
		return nil, err
	}

	for _, fd := range fds {
		var builder strings.Builder

		// Write syntax
		builder.WriteString(fmt.Sprintf("syntax = \"%s\";\n\n", fd.Syntax()))

		// Write package
		if fd.Package() != "" {
			builder.WriteString(fmt.Sprintf("package %s%s;\n\n", strings.Replace(prefix, "/", ".", 1), fd.Package()))
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
				builder.WriteString(fmt.Sprintf("import \"%s%s\";\n", prefix, importPath))
				hasImports = true
			}
		}
		if hasImports {
			builder.WriteString("\n")
		}

		// Write file options exactly as they are (no modifications)
		opts := fd.Options().(*descriptorpb.FileOptions)
		if opts != nil {
			if opts.GoPackage != nil {
				builder.WriteString(fmt.Sprintf("option go_package = \"%s\";\n", *opts.GoPackage))
			}
			if opts.JavaMultipleFiles != nil {
				builder.WriteString(fmt.Sprintf("option java_multiple_files = %t;\n", *opts.JavaMultipleFiles))
			}
			if opts.JavaOuterClassname != nil {
				builder.WriteString(fmt.Sprintf("option java_outer_classname = \"%s\";\n", *opts.JavaOuterClassname))
			}
			if opts.JavaPackage != nil {
				builder.WriteString(fmt.Sprintf("option java_package = \"%s\";\n", *opts.JavaPackage))
			}
		}
		builder.WriteString("\n")

		// Write enums
		enums := fd.Enums()
		for i := 0; i < enums.Len(); i++ {
			enum := enums.Get(i)
			writeEnum(&builder, enum)
		}

		// Write messages
		messages := fd.Messages()
		for i := 0; i < messages.Len(); i++ {
			msg := messages.Get(i)
			writeMessage(&builder, msg, 0)
		}

		// Write services - special handling for google/longrunning
		services := fd.Services()
		for i := 0; i < services.Len(); i++ {
			svc := services.Get(i)
			writeService(&builder, svc, fd)
		}

		// Store the built proto content in map with key as outputDir + fd.Path()
		key := filepath.Join(outputDir, fd.Path())
		fileDetails[key] = []byte(builder.String())
	}

	return fileDetails, nil
}

// writeEnum writes the declaration and values of an enum to the builder.
//
// builder: The strings.Builder to write to.
// enum: The EnumDescriptor representing the enum.
func writeEnum(builder *strings.Builder, enum protoreflect.EnumDescriptor) {
	builder.WriteString(fmt.Sprintf("enum %s {\n", enum.Name()))
	values := enum.Values()
	for i := 0; i < values.Len(); i++ {
		value := values.Get(i)
		builder.WriteString(fmt.Sprintf("  %s = %d;\n", value.Name(), value.Number()))
	}
	builder.WriteString("}\n\n")
}

// writeMessage writes the declaration, fields, nested enums, and nested messages of a message to the builder.
//
// builder: The strings.Builder to write to.
// msg: The MessageDescriptor representing the message.
// indent: Indentation level (number of indents to apply).
func writeMessage(builder *strings.Builder, msg protoreflect.MessageDescriptor, indent int) {
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
		fields := msg.Fields()
		for j := 0; j < fields.Len(); j++ {
			field := fields.Get(j)
			if field.ContainingOneof() == oneof {
				writeField(builder, field, indent+2)
			}
		}
		builder.WriteString(fmt.Sprintf("%s  }\n", indentStr))
	}

	// Write regular fields (not in oneofs)
	fields := msg.Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		if field.ContainingOneof() == nil {
			writeField(builder, field, indent+1)
		}
	}

	// Write nested messages
	nestedMessages := msg.Messages()
	for i := 0; i < nestedMessages.Len(); i++ {
		nested := nestedMessages.Get(i)
		writeMessage(builder, nested, indent+1)
	}

	builder.WriteString(fmt.Sprintf("%s}\n\n", indentStr))
}

// writeField writes a single field declaration to the builder, with appropriate type and indentation.
//
// builder: The strings.Builder to write to.
// field: The FieldDescriptor representing the field.
// indent: Indentation level (number of indents to apply).
func writeField(builder *strings.Builder, field protoreflect.FieldDescriptor, indent int) {
	indentStr := strings.Repeat("  ", indent)
	fieldType := getFieldType(field)
	builder.WriteString(fmt.Sprintf("%s%s %s = %d;\n",
		indentStr, fieldType, field.Name(), field.Number()))
}

// getFieldType returns the string representation of a field's type,
// including repeated, message, and enum types with fully qualified names if necessary.
//
// field: The FieldDescriptor to analyze.
// Returns the field type as a string.
func getFieldType(field protoreflect.FieldDescriptor) string {
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
		msgDesc := field.Message()
		if msgDesc.Parent() != nil {
			baseType = getFullMessageName(msgDesc)
		} else {
			if pkg := msgDesc.ParentFile().Package(); pkg != "" {
				baseType = string(pkg) + "." + string(msgDesc.Name())
			} else {
				baseType = string(msgDesc.Name())
			}
		}
	case protoreflect.EnumKind:
		enumDesc := field.Enum()
		if enumDesc.Parent() != nil {
			baseType = getFullEnumName(enumDesc)
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

// getFullMessageName returns the fully qualified name of a message, including nested paths and package.
//
// msgDesc: The MessageDescriptor to analyze.
// Returns the full message name as a string.
func getFullMessageName(msgDesc protoreflect.MessageDescriptor) string {
	var parts []string
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
	if pkg := msgDesc.ParentFile().Package(); pkg != "" {
		return string(pkg) + "." + strings.Join(parts, ".")
	}
	return strings.Join(parts, ".")
}

// getFullEnumName returns the fully qualified name of an enum, including nested paths and package.
//
// enumDesc: The EnumDescriptor to analyze.
// Returns the full enum name as a string.
func getFullEnumName(enumDesc protoreflect.EnumDescriptor) string {
	var parts []string
	parts = append(parts, string(enumDesc.Name()))
	if parent := enumDesc.Parent(); parent != nil {
		if parentMsg, ok := parent.(protoreflect.MessageDescriptor); ok {
			parentName := getFullMessageName(parentMsg)
			return parentName + "." + strings.Join(parts, ".")
		}
	}
	if pkg := enumDesc.ParentFile().Package(); pkg != "" {
		return string(pkg) + "." + strings.Join(parts, ".")
	}
	return strings.Join(parts, ".")
}

// writeService writes the declaration of a service and its methods to the builder.
//
// For services in the "google.longrunning" package, the service is not written.
// builder: The strings.Builder to write to.
// svc: The ServiceDescriptor representing the service.
// fd: The FileDescriptor containing the service.
func writeService(builder *strings.Builder, svc protoreflect.ServiceDescriptor, fd protoreflect.FileDescriptor) {
	// Special handling for google/longrunning - only write messages, not RPC calls
	if strings.Contains(string(fd.Package()), "google.longrunning") {
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
