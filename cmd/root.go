package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/machanirobotics/buffman/cmd/convert"
	"github.com/machanirobotics/buffman/cmd/generate"
	"github.com/machanirobotics/buffman/pkg/runner"
	"github.com/spf13/cobra"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "buffman",
	Short: "Buffman - a simple cli tool to convert proto files to flatbuffers, you can use a buffman.yml to do conversions and generations",
	Run: func(cmd *cobra.Command, args []string) {
		r := runner.NewRunner()
		if err := r.Run(context.Background(), configPath); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(convert.ConvertCmd, generate.GenerateCmd)
	rootCmd.Flags().StringVarP(&configPath, "file", "f", "buffman.yml", "file path of the buffman config")
}
