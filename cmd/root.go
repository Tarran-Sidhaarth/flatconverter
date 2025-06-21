package cmd

import (
	"os"

	"github.com/machanirobotics/buffman/cmd/convert"
	"github.com/machanirobotics/buffman/cmd/generate"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "buffman",
	Short: "Buffman - a simple cli tool to convert proto files to flatbuffers",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(convert.ConvertCmd, generate.GenerateCmd)
}
