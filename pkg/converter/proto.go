// // pkg/converter/buf_converter.go
package converter

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"strings"

// 	"github.com/bufbuild/protocompile"
// 	"github.com/bufbuild/protocompile/reporter"
// 	"google.golang.org/protobuf/reflect/protoreflect"
// )

// // BufConverter implements the Converter interface using bufbuild/protocompile
// type BufConverter struct {
// 	cleanedFiles map[string]*ProtoContent
// }

// // NewBufConverter creates a new converter instance
// func NewBufConverter() *BufConverter {
// 	return &BufConverter{
// 		cleanedFiles: make(map[string]*ProtoContent),
// 	}
// }

// func (b *BufConverter) Clean(ctx context.Context, manipulations []Manipulation, opts CleanOptions) error {
// 	// Find all proto files in the input directories
// 	allProtoFiles, err := b.findAllProtoFiles(opts.InputPaths)
// 	if err != nil {
// 		return fmt.Errorf("failed to find proto files: %w", err)
// 	}

// 	if len(allProtoFiles) == 0 {
// 		return fmt.Errorf("no proto files found in specified paths")
// 	}

// 	fmt.Printf("Found %d proto files to process\n", len(allProtoFiles))

// 	// Convert absolute paths to relative paths for compilation
// 	relativeProtoFiles := b.convertToRelativePaths(allProtoFiles, opts.ImportPaths)

// 	// Create compiler with proper import paths
// 	compiler := protocompile.Compiler{
// 		Resolver: &protocompile.SourceResolver{
// 			ImportPaths: opts.ImportPaths, // Use the base import paths
// 		},
// 		Reporter: reporter.NewReporter(nil, nil),
// 	}

// 	// Process files in batches using relative paths
// 	batchSize := 10
// 	for i := 0; i < len(relativeProtoFiles); i += batchSize {
// 		end := i + batchSize
// 		if end > len(relativeProtoFiles) {
// 			end = len(relativeProtoFiles)
// 		}

// 		batch := relativeProtoFiles[i:end]
// 		err := b.processBatch(ctx, compiler, batch, opts, manipulations, allProtoFiles[i:end])
// 		if err != nil {
// 			fmt.Printf("Warning: failed to process batch starting at %d: %v\n", i, err)
// 			continue
// 		}
// 	}

// 	fmt.Printf("Successfully processed %d proto files\n", len(b.cleanedFiles))
// 	return nil
// }

// func (b *BufConverter) findAllProtoFiles(inputPaths []string) ([]string, error) {
// 	var allFiles []string

// 	for _, inputPath := range inputPaths {
// 		err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
// 			if err != nil {
// 				return err
// 			}

// 			if !info.IsDir() && strings.HasSuffix(path, ".proto") {
// 				// Skip Google API proto files entirely
// 				if b.isGoogleAPIProtoFile(path) {
// 					fmt.Printf("Skipping Google API file: %s\n", path)
// 					return nil
// 				}

// 				allFiles = append(allFiles, path)
// 			}

// 			return nil
// 		})

// 		if err != nil {
// 			return nil, fmt.Errorf("failed to walk directory %s: %w", inputPath, err)
// 		}
// 	}

// 	return allFiles, nil
// }

// func (b *BufConverter) isGoogleAPIProtoFile(path string) bool {
// 	// Skip files in google/api directory or files that are Google API related
// 	return strings.Contains(path, "google/api/") ||
// 		strings.Contains(path, "google\\api\\") ||
// 		strings.HasSuffix(path, "field_behavior.proto") ||
// 		strings.HasSuffix(path, "resource.proto") ||
// 		strings.HasSuffix(path, "annotations.proto") ||
// 		strings.HasSuffix(path, "http.proto") ||
// 		strings.HasSuffix(path, "client.proto")
// }

// // Convert absolute file paths to relative paths for compilation
// func (b *BufConverter) convertToRelativePaths(absolutePaths []string, importPaths []string) []string {
// 	var relativePaths []string

// 	for _, absPath := range absolutePaths {
// 		// Find which import path this file belongs to
// 		var relativePath string
// 		for _, importPath := range importPaths {
// 			if strings.HasPrefix(absPath, importPath+string(filepath.Separator)) {
// 				// Remove the import path prefix to get relative path
// 				relativePath = strings.TrimPrefix(absPath, importPath+string(filepath.Separator))
// 				break
// 			} else if absPath == importPath {
// 				// File is directly in the import path
// 				relativePath = filepath.Base(absPath)
// 				break
// 			}
// 		}

// 		if relativePath == "" {
// 			// Fallback: use the absolute path
// 			relativePath = absPath
// 		}

// 		relativePaths = append(relativePaths, relativePath)
// 	}

// 	return relativePaths
// }

// // Updated processBatch to handle both relative and absolute paths
// func (b *BufConverter) processBatch(ctx context.Context, compiler protocompile.Compiler, relativeFiles []string, opts CleanOptions, manipulations []Manipulation, absoluteFiles []string) error {
// 	// Try to compile the batch using relative paths
// 	fds, err := compiler.Compile(ctx, relativeFiles...)
// 	if err != nil {
// 		// If batch fails, try individual files
// 		fmt.Printf("Batch compilation failed, trying individual files: %v\n", err)
// 		return b.processIndividualFiles(ctx, compiler, relativeFiles, opts, manipulations, absoluteFiles)
// 	}

// 	// Process each successfully compiled file descriptor
// 	for i, fd := range fds {
// 		// Use the absolute path for output file structure
// 		absolutePath := absoluteFiles[i]
// 		err := b.processFileDescriptorWithPath(fd, absolutePath, opts, manipulations)
// 		if err != nil {
// 			fmt.Printf("Warning: failed to process file %s: %v\n", fd.Path(), err)
// 			continue
// 		}
// 	}

// 	return nil
// }

// func (b *BufConverter) processIndividualFiles(ctx context.Context, compiler protocompile.Compiler, relativeFiles []string, opts CleanOptions, manipulations []Manipulation, absoluteFiles []string) error {
// 	for i, relativeFile := range relativeFiles {
// 		fds, err := compiler.Compile(ctx, relativeFile)
// 		if err != nil {
// 			fmt.Printf("Warning: failed to compile %s: %v\n", relativeFile, err)
// 			continue
// 		}

// 		for _, fd := range fds {
// 			absolutePath := absoluteFiles[i]
// 			err := b.processFileDescriptorWithPath(fd, absolutePath, opts, manipulations)
// 			if err != nil {
// 				fmt.Printf("Warning: failed to process file %s: %v\n", fd.Path(), err)
// 				continue
// 			}
// 		}
// 	}
// 	return nil
// }

// // Updated to use the absolute path for output structure
// func (b *BufConverter) processFileDescriptorWithPath(fd protoreflect.FileDescriptor, absolutePath string, opts CleanOptions, manipulations []Manipulation) error {
// 	// Reconstruct proto content from parsed descriptor
// 	content := &ProtoContent{
// 		Content:  b.reconstructProtoFromDescriptor(fd),
// 		FilePath: absolutePath, // Use absolute path for file operations
// 		Package:  string(fd.Package()),
// 		Messages: b.extractMessageInfo(fd),
// 		Services: b.extractServiceInfo(fd),
// 	}

// 	// Apply additional manipulations if any
// 	for _, manipulation := range manipulations {
// 		err := manipulation.Apply(content)
// 		if err != nil {
// 			return fmt.Errorf("failed to apply manipulation %s: %w",
// 				manipulation.Name(), err)
// 		}
// 	}

// 	// Store cleaned content
// 	b.cleanedFiles[content.FilePath] = content

// 	// Write cleaned file maintaining directory structure
// 	outputPath := b.getOutputPathWithStructure(content.FilePath, opts)
// 	err := b.writeFile(content.Content, outputPath)
// 	if err != nil {
// 		return fmt.Errorf("failed to write cleaned file: %w", err)
// 	}

// 	fmt.Printf("Processed: %s -> %s\n", content.FilePath, outputPath)
// 	return nil
// }

// func (b *BufConverter) getOutputPathWithStructure(inputPath string, opts CleanOptions) string {
// 	if !opts.CreateNewDirectory {
// 		return inputPath
// 	}

// 	// Find the base directory from input paths to maintain relative structure
// 	var relativePath string
// 	for _, inputDir := range opts.InputPaths {
// 		if strings.HasPrefix(inputPath, inputDir) {
// 			relativePath = strings.TrimPrefix(inputPath, inputDir)
// 			relativePath = strings.TrimPrefix(relativePath, string(filepath.Separator))
// 			break
// 		}
// 	}

// 	if relativePath == "" {
// 		// Fallback to just the filename
// 		relativePath = filepath.Base(inputPath)
// 	}

// 	return filepath.Join(opts.OutputDirectory, relativePath)
// }

// func (b *BufConverter) reconstructProtoFromDescriptor(fd protoreflect.FileDescriptor) string {
// 	var builder strings.Builder

// 	// Write syntax
// 	builder.WriteString(fmt.Sprintf("syntax = \"%s\";\n\n", fd.Syntax()))

// 	// Write package
// 	if fd.Package() != "" {
// 		builder.WriteString(fmt.Sprintf("package %s;\n\n", fd.Package()))
// 	}

// 	// Write clean imports (excluding Google API imports)
// 	b.writeCleanImports(&builder, fd)

// 	// Write file options with renamed packages
// 	b.writeRenamedFileOptions(&builder, fd)

// 	// Write messages
// 	messages := fd.Messages()
// 	for i := 0; i < messages.Len(); i++ {
// 		b.writeMessageFromDescriptor(&builder, messages.Get(i), 0)
// 	}

// 	// Write services
// 	services := fd.Services()
// 	for i := 0; i < services.Len(); i++ {
// 		b.writeServiceFromDescriptor(&builder, services.Get(i))
// 	}

// 	return builder.String()
// }

// func (b *BufConverter) writeCleanImports(builder *strings.Builder, fd protoreflect.FileDescriptor) {
// 	imports := fd.Imports()
// 	hasNonGoogleImports := false

// 	for i := 0; i < imports.Len(); i++ {
// 		imp := imports.Get(i)
// 		importPath := string(imp.Path())

// 		// Skip Google API imports
// 		if !b.isGoogleAPIImport(importPath) {
// 			builder.WriteString(fmt.Sprintf("import \"%s\";\n", importPath))
// 			hasNonGoogleImports = true
// 		}
// 	}

// 	if hasNonGoogleImports {
// 		builder.WriteString("\n")
// 	}
// }

// func (b *BufConverter) writeRenamedFileOptions(builder *strings.Builder, fd protoreflect.FileDescriptor) {
// 	// Extract package name to construct renamed options
// 	packageName := string(fd.Package())

// 	// Only write options for machanirobotics packages
// 	if !strings.Contains(packageName, "machanirobotics") {
// 		return
// 	}

// 	// Convert package path for Go and Java options
// 	packageParts := strings.Split(packageName, ".")
// 	var goPackageSuffix string
// 	var javaPackagePath string

// 	if len(packageParts) >= 2 {
// 		// Get the last two parts for Go package suffix
// 		goPackageSuffix = packageParts[len(packageParts)-2] + "pbv1"

// 		// Reconstruct Java package with "fb" inserted
// 		javaPackagePath = strings.Join(packageParts, ".")
// 		javaPackagePath = strings.Replace(javaPackagePath, ".engine.", ".fb.engine.", 1)
// 	}

// 	// Write renamed Go package option
// 	builder.WriteString(fmt.Sprintf("option go_package = \"github.com/machanirobotics/protoverse-go/fb/engine/%s;%s\";\n",
// 		goPackageSuffix, goPackageSuffix))

// 	// Write standard Java options
// 	builder.WriteString("option java_multiple_files = true;\n")
// 	builder.WriteString("option java_outer_classname = \"ConfigProto\";\n")

// 	// Write renamed Java package option
// 	if javaPackagePath != "" {
// 		builder.WriteString(fmt.Sprintf("option java_package = \"%s\";\n", javaPackagePath))
// 	}

// 	builder.WriteString("\n")
// }

// func (b *BufConverter) writeMessageFromDescriptor(builder *strings.Builder, msg protoreflect.MessageDescriptor, indent int) {
// 	indentStr := strings.Repeat("  ", indent)

// 	// Write message comment
// 	builder.WriteString(fmt.Sprintf("%s// %s configuration\n", indentStr, msg.Name()))
// 	builder.WriteString(fmt.Sprintf("%smessage %s {\n", indentStr, msg.Name()))

// 	// Write fields from descriptor
// 	fields := msg.Fields()
// 	for i := 0; i < fields.Len(); i++ {
// 		field := fields.Get(i)
// 		b.writeFieldFromDescriptor(builder, field, indent+1)
// 	}

// 	// Write nested messages
// 	nestedMessages := msg.Messages()
// 	for i := 0; i < nestedMessages.Len(); i++ {
// 		b.writeMessageFromDescriptor(builder, nestedMessages.Get(i), indent+1)
// 	}

// 	builder.WriteString(fmt.Sprintf("%s}\n\n", indentStr))
// }

// func (b *BufConverter) writeFieldFromDescriptor(builder *strings.Builder, field protoreflect.FieldDescriptor, indent int) {
// 	indentStr := strings.Repeat("  ", indent)

// 	// Write field comment
// 	fieldComment := b.generateFieldComment(string(field.Name()))
// 	builder.WriteString(fmt.Sprintf("%s// %s\n", indentStr, fieldComment))

// 	// Get field type from descriptor
// 	fieldType := b.getFieldTypeFromDescriptor(field)

// 	// Write field definition
// 	builder.WriteString(fmt.Sprintf("%s%s %s = %d;\n",
// 		indentStr, fieldType, field.Name(), field.Number()))
// }

// func (b *BufConverter) writeServiceFromDescriptor(builder *strings.Builder, svc protoreflect.ServiceDescriptor) {
// 	builder.WriteString(fmt.Sprintf("service %s {\n", svc.Name()))

// 	methods := svc.Methods()
// 	for i := 0; i < methods.Len(); i++ {
// 		method := methods.Get(i)
// 		builder.WriteString(fmt.Sprintf("  rpc %s(%s) returns (%s);\n",
// 			method.Name(),
// 			method.Input().Name(),
// 			method.Output().Name()))
// 	}

// 	builder.WriteString("}\n\n")
// }

// func (b *BufConverter) getFieldTypeFromDescriptor(field protoreflect.FieldDescriptor) string {
// 	prefix := ""
// 	if field.Cardinality() == protoreflect.Repeated {
// 		prefix = "repeated "
// 	}

// 	var baseType string
// 	switch field.Kind() {
// 	case protoreflect.StringKind:
// 		baseType = "string"
// 	case protoreflect.Int32Kind:
// 		baseType = "int32"
// 	case protoreflect.Int64Kind:
// 		baseType = "int64"
// 	case protoreflect.Uint32Kind:
// 		baseType = "uint32"
// 	case protoreflect.Uint64Kind:
// 		baseType = "uint64"
// 	case protoreflect.BoolKind:
// 		baseType = "bool"
// 	case protoreflect.FloatKind:
// 		baseType = "float"
// 	case protoreflect.DoubleKind:
// 		baseType = "double"
// 	case protoreflect.BytesKind:
// 		baseType = "bytes"
// 	case protoreflect.MessageKind:
// 		baseType = string(field.Message().Name())
// 	case protoreflect.EnumKind:
// 		baseType = string(field.Enum().Name())
// 	default:
// 		baseType = field.Kind().String()
// 	}

// 	return prefix + baseType
// }

// func (b *BufConverter) extractMessageInfo(fd protoreflect.FileDescriptor) []MessageInfo {
// 	var messages []MessageInfo

// 	msgDescs := fd.Messages()
// 	for i := 0; i < msgDescs.Len(); i++ {
// 		msg := msgDescs.Get(i)

// 		var fields []FieldInfo
// 		fieldDescs := msg.Fields()
// 		for j := 0; j < fieldDescs.Len(); j++ {
// 			field := fieldDescs.Get(j)
// 			fields = append(fields, FieldInfo{
// 				Name:     string(field.Name()),
// 				Type:     b.getFieldTypeFromDescriptor(field),
// 				Number:   int(field.Number()),
// 				Repeated: field.Cardinality() == protoreflect.Repeated,
// 			})
// 		}

// 		messages = append(messages, MessageInfo{
// 			Name:   string(msg.Name()),
// 			Fields: fields,
// 		})
// 	}

// 	return messages
// }

// func (b *BufConverter) extractServiceInfo(fd protoreflect.FileDescriptor) []ServiceInfo {
// 	var services []ServiceInfo

// 	svcDescs := fd.Services()
// 	for i := 0; i < svcDescs.Len(); i++ {
// 		svc := svcDescs.Get(i)

// 		var methods []MethodInfo
// 		methodDescs := svc.Methods()
// 		for j := 0; j < methodDescs.Len(); j++ {
// 			method := methodDescs.Get(j)
// 			methods = append(methods, MethodInfo{
// 				Name:       string(method.Name()),
// 				InputType:  string(method.Input().Name()),
// 				OutputType: string(method.Output().Name()),
// 			})
// 		}

// 		services = append(services, ServiceInfo{
// 			Name:    string(svc.Name()),
// 			Methods: methods,
// 		})
// 	}

// 	return services
// }

// func (b *BufConverter) Convert(ctx context.Context, opts ConvertOptions) error {
// 	if len(b.cleanedFiles) == 0 {
// 		return fmt.Errorf("no cleaned files available for conversion. Run Clean() first")
// 	}

// 	switch opts.OutputFormat {
// 	case OutputFormatFlatBuffer:
// 		return b.convertToFlatBufferWithFlatc(ctx, opts)
// 	default:
// 		return fmt.Errorf("unsupported output format: %s", opts.OutputFormat)
// 	}
// }

// func (b *BufConverter) convertToFlatBufferWithFlatc(ctx context.Context, opts ConvertOptions) error {
// 	// Create output directory
// 	err := os.MkdirAll(opts.OutputDirectory, 0o755)
// 	if err != nil {
// 		return fmt.Errorf("failed to create FlatBuffer output directory: %w", err)
// 	}

// 	// Get the cleaned proto files directory
// 	cleanedDir := ""
// 	for filePath := range b.cleanedFiles {
// 		dir := filepath.Dir(filePath)
// 		if cleanedDir == "" || len(dir) < len(cleanedDir) {
// 			cleanedDir = dir
// 		}
// 	}

// 	// Find the root directory of cleaned files
// 	if cleanedDir == "" {
// 		return fmt.Errorf("no cleaned files found")
// 	}

// 	// Use flatc to convert all proto files
// 	for filePath := range b.cleanedFiles {
// 		err := b.convertSingleFileWithFlatc(filePath, opts.OutputDirectory, cleanedDir)
// 		if err != nil {
// 			fmt.Printf("Warning: failed to convert %s with flatc: %v\n", filePath, err)
// 			continue
// 		}
// 		fmt.Printf("Converted: %s\n", filePath)
// 	}

// 	return nil
// }

// func (b *BufConverter) convertSingleFileWithFlatc(protoFile, targetDir, cleanedDir string) error {
// 	// Build flatc command
// 	cmd := exec.Command("flatc",
// 		"--proto",        // Convert from proto
// 		"-I", cleanedDir, // Include directory for imports
// 		"-o", targetDir, // Output directory
// 		protoFile, // Input proto file
// 	)

// 	// Set working directory
// 	cmd.Dir = "."

// 	// Capture output for debugging
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return fmt.Errorf("flatc command failed: %w\nOutput: %s", err, string(output))
// 	}

// 	return nil
// }

// // Helper methods
// func (b *BufConverter) isGoogleAPIImport(path string) bool {
// 	googleAPIImports := []string{
// 		"google/api/field_behavior.proto",
// 		"google/api/resource.proto",
// 		"google/api/annotations.proto",
// 		"google/api/http.proto",
// 		"google/api/client.proto",
// 	}

// 	for _, apiImport := range googleAPIImports {
// 		if path == apiImport {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (b *BufConverter) generateFieldComment(fieldName string) string {
// 	return fmt.Sprintf("the %s", strings.ReplaceAll(fieldName, "_", " "))
// }

// func (b *BufConverter) writeFile(content, path string) error {
// 	dir := filepath.Dir(path)
// 	err := os.MkdirAll(dir, 0o755)
// 	if err != nil {
// 		return fmt.Errorf("failed to create directory: %w", err)
// 	}

// 	err = os.WriteFile(path, []byte(content), 0o644)
// 	if err != nil {
// 		return fmt.Errorf("failed to write file: %w", err)
// 	}

// 	return nil
// }
