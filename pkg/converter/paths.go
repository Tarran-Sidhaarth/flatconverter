// Package converter provides utilities for converting Protocol Buffer files
// and managing file operations during the conversion process.
package converter

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// listProtoFiles returns a list of .proto files under rootPath, excluding
// google/api files and google/protobuf/descriptor.proto.
//
// The returned file paths are relative to rootPath. For example, if rootPath
// is "/home/xyz" and a file is "/home/xyz/sample.proto", the returned path
// will be "sample.proto".
//
// Directories are ignored. Only files with a ".proto" extension are included.
// Returns an error if directory traversal fails.
func listProtoFiles(rootPath string) ([]string, error) {
	var paths []string
	err := filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		// Skip Google API files and descriptor.proto
		if strings.Contains(path, "google/api") ||
			strings.Contains(path, "google/protobuf/descriptor.proto") {
			return nil
		}
		if strings.HasSuffix(path, ".proto") {
			relativePath := strings.TrimPrefix(path, rootPath)
			relativePath = strings.TrimPrefix(relativePath, string(filepath.Separator))
			paths = append(paths, relativePath)
		}
		return nil
	})
	return paths, err
}

// generateFiles writes files to disk based on the provided fileDetails map.
//
// fileDetails is a map where the key is the target file path and the value is
// the file content as a byte slice. The function creates any necessary
// directories for each file and writes the file content to disk.
//
// Returns an error if any directory or file cannot be created or written.
func generateFiles(fileDetails map[string][]byte) error {
	for path, content := range fileDetails {
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
		if err := os.WriteFile(path, content, 0o644); err != nil {
			return fmt.Errorf("failed to write file %s: %v", path, err)
		}
		fmt.Printf("Created: %s\n", path)
	}
	return nil
}
