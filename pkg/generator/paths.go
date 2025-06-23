package generator

import (
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
