package generator

import (
	"fmt"
	"os/exec"
	"strings"
)

const flatcCommand = "flatc %s %s -o %s %s %s" // include paths, language, output directory, package prefix, fbs file directory

func executeCommand(cmdStr string) error {
	// Split the command string into command and arguments
	cmdArgs := strings.Fields(cmdStr)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	// Run the command and capture output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute flatc command: %v, output: %s", err, string(output))
	}
	return nil
}
