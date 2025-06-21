package convert

import "github.com/spf13/cobra"

var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "this command is used to convert the proto files to desired format",
}

func init() {
	ConvertCmd.AddCommand(flatbuffersCmd)
}
