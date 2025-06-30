// Package converter implements conversion of Protocol Buffer files to FlatBuffers schemas.
package flatbuffers

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bufbuild/protocompile"
	"github.com/bufbuild/protocompile/reporter"
	"github.com/machanirobotics/buffman/internal/options"
	"github.com/machanirobotics/buffman/internal/template"
	"github.com/machanirobotics/buffman/internal/utilities"
)

// FlatbuffersParser handles conversion of Protocol Buffer files to FlatBuffers schemas.
//
// It manages compilation, cleaning, and file structure for the conversion process.
type FlatbuffersParser struct {
	compiler      *protocompile.Compiler
	protoDir      string
	cleanedDir    string
	flatbufferDir string
}

// NewFlatConverter returns a new FlatbuffersParser for converting proto files to FlatBuffers schemas.
//
// protoDir is the directory containing the input .proto files.
// cleanedDir is the directory where cleaned proto files (without google API imports) will be written.
// targetDir is the output directory for generated FlatBuffers schemas.
func NewFlatbuffersParser() (*FlatbuffersParser, error) {
	tempDir, err := os.MkdirTemp("", "cleaned")
	if err != nil {
		return nil, err
	}
	return &FlatbuffersParser{
		cleanedDir: tempDir,
	}, nil
}

// Convert performs the full conversion process from proto files to FlatBuffers schemas.
//
// It first removes google API imports from proto files, then generates FlatBuffers
// schemas using the flatc compiler. If keepCleaned is false, the cleaned proto files
// are deleted after conversion.
//
// Returns an error if any step of the process fails.
func (c *FlatbuffersParser) Parse(ctx context.Context, opts options.ParseOptions) error {
	c.compiler = &protocompile.Compiler{
		Resolver: &protocompile.SourceResolver{
			ImportPaths: []string{opts.InputDir},
		},
		Reporter: reporter.NewReporter(nil, nil),
	}
	c.protoDir = opts.InputDir
	c.flatbufferDir = opts.OutputDir

	if err := c.clearGoogleAPI(ctx); err != nil {
		return fmt.Errorf("could not remove google api from protos %v", err)
	}
	if err := os.MkdirAll(c.flatbufferDir, 0755); err != nil {
		return fmt.Errorf("failed to create flatbuffers directory: %v", err)
	}
	protoFiles, err := utilities.ListFilesRelativeToRoot(c.cleanedDir, ".proto")
	if err != nil {
		return fmt.Errorf("failed to find proto files: %v", err)
	}
	if len(protoFiles) == 0 {
		return errors.New("no proto files found")
	}
	t := template.NewTemplate("//")
	comment := t.BuildDefaultComment("Flatbuffers")

	fmt.Printf("🔄 Generating FlatBuffers schemas...\n")
	fmt.Printf("   Source: %s\n", c.cleanedDir)
	fmt.Printf("   Target: %s\n", c.flatbufferDir)
	fmt.Printf("   Files to process: %d\n\n", len(protoFiles))

	successCount := 0
	errorCount := 0

	for _, protoFile := range protoFiles {
		if err := c.convertProtoFile(ctx, protoFile); err != nil {
			fmt.Printf("✗ Failed to convert %s: %v\n", protoFile, err)
			errorCount++
		} else {
			if err := utilities.InsertGeneratedComments(comment, path.Join(c.flatbufferDir, strings.Replace(protoFile, ".proto", ".fbs", -1))); err != nil {
				fmt.Printf("✗ Failed to add generated comment: %v", err)
				errorCount++
			} else {
				fmt.Printf("✓ Converted: %s\n", protoFile)
				successCount++
			}
		}
	}

	fmt.Printf("\n📊 Conversion Summary:\n")
	fmt.Printf("   Successful: %d files\n", successCount)
	fmt.Printf("   Failed: %d files\n", errorCount)
	fmt.Printf("   Output directory: %s\n", c.flatbufferDir)

	if errorCount > 0 {
		return fmt.Errorf("conversion completed with %d errors", errorCount)
	}

	return os.RemoveAll(c.cleanedDir)
}

// fixFBSIncludes modifies include/import statements in a FlatBuffers schema file.
//
// It replaces import statements with include statements and changes .proto extensions
// to .fbs. If the file does not exist or cannot be written, an error is returned.
func (c *FlatbuffersParser) fixFBSIncludes(fbsFile string) error {
	if _, err := os.Stat(fbsFile); os.IsNotExist(err) {
		return fmt.Errorf("FBS file not found: %s", fbsFile)
	}
	content, err := os.ReadFile(fbsFile)
	if err != nil {
		return fmt.Errorf("failed to read FBS file: %w", err)
	}
	originalContent := string(content)
	importPattern := regexp.MustCompile(`import\s+"([^"]+)\.proto"\s*;`)
	modifiedContent := importPattern.ReplaceAllString(originalContent, `include "$1.fbs";`)
	includePattern := regexp.MustCompile(`include\s+"([^"]+)\.proto"`)
	modifiedContent = includePattern.ReplaceAllString(modifiedContent, `include "$1.fbs"`)
	if modifiedContent != originalContent {
		if err := os.WriteFile(fbsFile, []byte(modifiedContent), 0644); err != nil {
			return fmt.Errorf("failed to write modified FBS file: %w", err)
		}
		fmt.Printf("↳ Fixed includes in: %s\n", filepath.Base(fbsFile))
	}
	return nil
}
