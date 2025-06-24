package generator

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func listLanguageFiles(rootPath, extenstion string) ([]string, error) {
	var paths []string
	err := filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, extenstion) {
			relativePath := strings.TrimPrefix(path, rootPath)
			relativePath = strings.TrimPrefix(relativePath, string(filepath.Separator))
			paths = append(paths, relativePath)
		}
		return nil
	})
	return paths, err
}

// getIncludePaths walks the entire directory tree starting from the generator's
// base flatbuffer directory. It collects all subdirectory paths and formats them
// as `-I <path>` strings. This allows the flatc compiler to resolve imports
// between .fbs files located in different subdirectories.
// It returns a slice of formatted include path strings or an error if the
// directory walk fails.
func (f *FlatGenerator) getIncludePaths() ([]string, error) {
	var paths []string
	err := filepath.WalkDir(f.flatbufferDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Add every directory encountered to the include paths.
		if d.IsDir() {
			paths = append(paths, "-I "+path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk include directories: %w", err)
	}
	return paths, nil
}
