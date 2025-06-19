package generator

import (
	"fmt"
	"os/exec"
)

const flatcCommand = "flatc %s %s -o %s %s %s" // include paths, language, output directory, package prefix, fbs file directory

func executeCommand(cmdStr string) error {
	fullCmd := fmt.Sprintf("shopt -s globstar; %s", cmdStr)

	// Use bash -c to enable shell glob expansion
	cmd := exec.Command("bash", "-c", fullCmd)

	// Run the command and capture output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute flatc command: %v, output: %s", err, string(output))
	}
	return nil
}
