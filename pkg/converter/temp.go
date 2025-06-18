package converter

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func (c *Converter) Convert(ctx context.Context, flatbufferDir string) error {
	// Create the flatbuffers output directory
	if err := os.MkdirAll(flatbufferDir, 0755); err != nil {
		return fmt.Errorf("failed to create flatbuffers directory: %w", err)
	}

	// Find all proto files in the cleaned directory
	protoFiles, err := c.findProtoFiles()
	if err != nil {
		return fmt.Errorf("failed to find proto files: %w", err)
	}

	if len(protoFiles) == 0 {
		return fmt.Errorf("no proto files found in %s", c.DestinationDir)
	}

	fmt.Printf("🔄 Generating FlatBuffers schemas...\n")
	fmt.Printf("   Source: %s\n", c.DestinationDir)
	fmt.Printf("   Target: %s\n", flatbufferDir)
	fmt.Printf("   Files to process: %d\n\n", len(protoFiles))

	successCount := 0
	errorCount := 0

	for _, protoFile := range protoFiles {
		if err := c.convertProtoFile(ctx, protoFile, flatbufferDir); err != nil {
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
	fmt.Printf("   Output directory: %s\n", flatbufferDir)

	if errorCount > 0 {
		return fmt.Errorf("conversion completed with %d errors", errorCount)
	}

	return nil
}

func (c *Converter) findProtoFiles() ([]string, error) {
	var protoFiles []string

	err := filepath.Walk(c.DestinationDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".proto") {
			protoFiles = append(protoFiles, path)
		}

		return nil
	})

	return protoFiles, err
}

func (c *Converter) convertProtoFile(ctx context.Context, protoFile, flatbufferDir string) error {
	// Get relative path from cleaned directory
	relPath, err := filepath.Rel(c.DestinationDir, protoFile)
	if err != nil {
		return fmt.Errorf("failed to get relative path: %w", err)
	}

	// Create target directory structure in flatbuffers directory
	targetDir := filepath.Join(flatbufferDir, filepath.Dir(relPath))
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create target directory %s: %w", targetDir, err)
	}

	// Execute flatc command with proper include path
	// flatc --proto -I <include_path> -o <output_dir> <proto_file>
	cmd := exec.CommandContext(ctx, "flatc",
		"--proto",
		"-I", c.DestinationDir, // Include path for resolving imports
		"-o", targetDir, // Output directory
		protoFile, // Input proto file
	)

	// Capture both stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("flatc command failed: %w\nOutput: %s", err, string(output))
	}

	// Post-process the generated FBS file to fix includes
	fbsFileName := strings.TrimSuffix(filepath.Base(protoFile), ".proto") + ".fbs"
	fbsFilePath := filepath.Join(targetDir, fbsFileName)

	if err := c.fixFBSIncludes(fbsFilePath); err != nil {
		// Log warning but don't fail the conversion
		fmt.Printf("  ⚠️  Warning: failed to fix includes in %s: %v\n", fbsFileName, err)
	}

	return nil
}

func (c *Converter) fixFBSIncludes(fbsFile string) error {
	// Check if file exists
	if _, err := os.Stat(fbsFile); os.IsNotExist(err) {
		return fmt.Errorf("FBS file not found: %s", fbsFile)
	}

	// Read the file content
	content, err := os.ReadFile(fbsFile)
	if err != nil {
		return fmt.Errorf("failed to read FBS file: %w", err)
	}

	originalContent := string(content)

	// Convert import statements to include statements and change .proto to .fbs
	// Pattern: import "folder1/file.proto"; -> include "folder1/file.fbs";
	importPattern := regexp.MustCompile(`import\s+"([^"]+)\.proto"\s*;`)
	modifiedContent := importPattern.ReplaceAllString(originalContent, `include "$1.fbs";`)

	// Also handle include statements that might have .proto extensions
	includePattern := regexp.MustCompile(`include\s+"([^"]+)\.proto"`)
	modifiedContent = includePattern.ReplaceAllString(modifiedContent, `include "$1.fbs"`)

	// Write back if changes were made
	if modifiedContent != originalContent {
		if err := os.WriteFile(fbsFile, []byte(modifiedContent), 0644); err != nil {
			return fmt.Errorf("failed to write modified FBS file: %w", err)
		}
		fmt.Printf("  ↳ Fixed includes in: %s\n", filepath.Base(fbsFile))
	}

	return nil
}
