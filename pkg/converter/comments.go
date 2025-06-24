package converter

import (
	"os"
	"strings"

	"github.com/machanirobotics/buffman/pkg/template"
)

var t = template.NewTemplate("//")
var comment = t.BuildDefaultComment("Flatbuffers")

func (f *FlatConverter) insertGeneratedComments(fbsFilePath string) error {
	// Read the entire file content
	content, err := os.ReadFile(fbsFilePath)
	if err != nil {
		return err
	}

	// Convert to string and split into lines
	lines := strings.Split(string(content), "\n")

	// Remove the first line if there are any lines
	if len(lines) > 0 {
		lines = lines[1:]
	}

	// Prepare the comment to prepend
	commentStr := comment

	// Join remaining lines back together
	remainingContent := strings.Join(lines, "\n")

	// Create new content: comment + remaining content (without first line)
	newContent := commentStr + remainingContent

	// Write the new content back to the file
	err = os.WriteFile(fbsFilePath, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
