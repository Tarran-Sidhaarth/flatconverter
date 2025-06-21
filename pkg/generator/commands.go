package generator

import (
	"fmt"
	"os/exec"
)

// flatcCommand is a template for executing the FlatBuffers compiler (flatc).
// The format placeholders are intended for the following arguments:
// 1. Include paths for interdependent .fbs files (`-I`).
// 2. Language-specific generation flags (e.g., --go, --java).
// 3. Output directory path (`-o`).
// 4. Language-specific options (e.g., --go-pkg-prefix).
// 5. The path to the source .fbs files, which supports globbing (e.g., "path/**/*.fbs").
const flatcCommand = "flatc %s %s -o %s %s %s"

// executeCommand runs a command string within a bash shell.
// It enables the `globstar` shell option, allowing for recursive file
// matching (e.g., `**/*.fbs`), which is necessary for processing nested
// directories of schema files.
//
// It returns an error if the command fails, wrapping the original error
// with the combined stdout and stderr from the command for easier debugging.
func executeCommand(cmdStr string) error {
	// Prepend 'shopt -s globstar;' to the command to enable recursive globbing
	// for the duration of this command's execution.
	fullCmd := fmt.Sprintf("shopt -s globstar; %s", cmdStr)

	// Use "bash -c" to execute the command in a bash shell, which allows for the
	// interpretation of shell-specific syntax like globstar.
	cmd := exec.Command("bash", "-c", fullCmd)

	// Run the command and capture its combined standard output and standard error.
	output, err := cmd.CombinedOutput()
	if err != nil {
		// If the command returns a non-zero exit code, return an error that includes
		// the captured output to provide context for what went wrong.
		return fmt.Errorf("failed to execute command: %v\nOutput: %s", err, string(output))
	}

	return nil
}
