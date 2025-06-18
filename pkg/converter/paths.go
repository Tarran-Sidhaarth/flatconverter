package converter

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// listProtoFiles is a function that returns the list of protofiles in a given directory
// it ignores google/api files and google/protobuf/descriptor.proto
// it strips the prefix/root while returning, for example if the root is
// /home/xyz and the file is /home/xyz/sample.proto it returns sample.proto
func listProtoFiles(rootPath string) ([]string, error) {
	var paths []string
	err := filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// skipping directories from being included in the path
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
			relativePath := strings.TrimPrefix(path, rootPath)
			relativePath = strings.TrimPrefix(relativePath, string(filepath.Separator))
			paths = append(paths, relativePath)
		}
		return nil
	})
	return paths, err
}

// generateFiles is a function that takes in fileDetails where the key is the path and
// the value is the acual file content
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
