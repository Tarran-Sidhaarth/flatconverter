package generate

import "github.com/spf13/cobra"

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "this command is used to generate the language specific files",
}

func init() {
	GenerateCmd.AddCommand(flatbuffersCmd)
}
