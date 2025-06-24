package internal

import (
	"os/exec"
	"regexp"
)

const Version = "1.0.0"

func GetFlatCVersion() string {
	cmd := exec.Command("bash", "-c", "flatc --version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "unknown"
	}
	re := regexp.MustCompile(`\d+\.\d+\.\d+`)
	return re.FindString(string(output))
}
