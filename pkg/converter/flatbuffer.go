// Package converter implements conversion of Protocol Buffer files to FlatBuffers schemas.
package converter

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bufbuild/protocompile"
	"github.com/bufbuild/protocompile/reporter"
)

// FlatConverter handles conversion of Protocol Buffer files to FlatBuffers schemas.
//
// It manages compilation, cleaning, and file structure for the conversion process.
type FlatConverter struct {
	compiler      *protocompile.Compiler
	protoDir      string
	cleanedDir    string
	flatbufferDir string
	prefix        string
}

// NewFlatConverter returns a new FlatConverter for converting proto files to FlatBuffers schemas.
//
// protoDir is the directory containing the input .proto files.
// cleanedDir is the directory where cleaned proto files (without google API imports) will be written.
// targetDir is the output directory for generated FlatBuffers schemas.
func newFlatConverter(protoDir, targetDir, prefix string) (*FlatConverter, error) {
	tempDir, err := os.MkdirTemp("", "cleaned")
	if err != nil {
		return nil, err
	}
	return &FlatConverter{
		compiler: &protocompile.Compiler{
			Resolver: &protocompile.SourceResolver{
				ImportPaths: []string{protoDir},
			},
			Reporter: reporter.NewReporter(nil, nil),
		},
		protoDir:      protoDir,
		cleanedDir:    tempDir,
		flatbufferDir: targetDir,
		prefix:        prefix,
	}, nil
}

// Convert performs the full conversion process from proto files to FlatBuffers schemas.
//
// It first removes google API imports from proto files, then generates FlatBuffers
// schemas using the flatc compiler. If keepCleaned is false, the cleaned proto files
// are deleted after conversion.
//
// Returns an error if any step of the process fails.
func (c *FlatConverter) Convert(ctx context.Context) error {
	if err := c.removeGoogleAPI(ctx); err != nil {
		return fmt.Errorf("could not remove google api from protos %v", err)
	}
	if err := os.MkdirAll(c.flatbufferDir, 0755); err != nil {
		return fmt.Errorf("failed to create flatbuffers directory: %v", err)
	}
	protoFiles, err := listProtoFiles(c.cleanedDir)
	if err != nil {
		return fmt.Errorf("failed to find proto files: %v", err)
	}
	if len(protoFiles) == 0 {
		return fmt.Errorf("no proto files found in %s", c.cleanedDir)
	}

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
			fmt.Printf("✓ Converted: %s\n", protoFile)
			successCount++
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

// removeGoogleAPI removes google API imports from proto files and writes cleaned files to the cleanedDir.
//
// Returns an error if proto files cannot be listed, cleaned, or written.
func (c *FlatConverter) removeGoogleAPI(ctx context.Context) error {
	cleanedDir := c.cleanedDir
	files, err := listProtoFiles(c.protoDir)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("no proto files found")
	}
	if c.prefix != "" {
		cleanedDir = path.Join(c.cleanedDir, c.prefix)
	}
	fileDetails, err := removeGoogleAPI(ctx, c.compiler, files, cleanedDir, c.prefix)
	if err != nil {
		return err
	}

	// Create files and necessary directories
	return generateFiles(fileDetails)
}

// convertProtoFile converts a single proto file to a FlatBuffers schema using flatc.
//
// It creates the necessary output directory structure and post-processes the generated
// .fbs file to fix include statements.
//
// Returns an error if conversion or post-processing fails.
func (c *FlatConverter) convertProtoFile(ctx context.Context, protoFile string) error {
	relPath, err := filepath.Rel(c.cleanedDir, path.Join(c.cleanedDir, protoFile))
	if err != nil {
		return fmt.Errorf("failed to get relative path: %w", err)
	}
	targetDir := filepath.Join(c.flatbufferDir, filepath.Dir(relPath))
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create target directory %s: %w", targetDir, err)
	}
	cmd := exec.CommandContext(ctx, "flatc",
		"--proto",
		"-I", c.cleanedDir,
		"-o", targetDir,
		path.Join(c.cleanedDir, protoFile),
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("flatc command failed: %w\nOutput: %s", err, string(output))
	}
	fbsFileName := strings.TrimSuffix(filepath.Base(protoFile), ".proto") + ".fbs"
	fbsFilePath := filepath.Join(targetDir, fbsFileName)
	if err := c.fixFBSIncludes(fbsFilePath); err != nil {
		fmt.Printf("  ⚠️  Warning: failed to fix includes in %s: %v\n", fbsFileName, err)
	}
	return nil
}

// fixFBSIncludes modifies include/import statements in a FlatBuffers schema file.
//
// It replaces import statements with include statements and changes .proto extensions
// to .fbs. If the file does not exist or cannot be written, an error is returned.
func (c *FlatConverter) fixFBSIncludes(fbsFile string) error {
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
