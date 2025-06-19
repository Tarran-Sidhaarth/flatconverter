package cmd

import (
	"os"

	"github.com/machanirobotics/flatconverter/cmd/flat"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "protox",
	Short: "Protox - a simple cli tool to convert proto files to flatbuffers",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(flat.FlatCmd)
}
