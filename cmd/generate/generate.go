// package generate implements the "generate" command and its subcommands
// for the Buffman CLI. It serves as a container for different code
// generation functionalities.
package generate

import (
	"context"
	"fmt"
	"os"

	"github.com/machanirobotics/buffman/internal/runner"
	"github.com/spf13/cobra"
)

var (
	buffmanConfigPath string
)

// GenerateCmd represents the base command for the "generate" functionality.
// It groups all code generation subcommands, such as 'flatbuffers'.
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Contains subcommands for generating code from schema files",
	Long: `The 'generate' command provides access to various code generators.

You can use this command in two ways:

1. With a config file:
   Use the -f flag to specify a buffman config file that contains all generation settings.

2. With subcommands:
   Use a subcommand to specify the type of code to generate. For example, to
   generate language-specific files from FlatBuffer schemas, use the 'flatbuffers'
   subcommand.

Examples:
  # Using config file
  buffman generate -f ./buffman.yaml

  # Using subcommand
  buffman generate flatbuffers --language=go --target_dir=./gen`,
	Run: func(cmd *cobra.Command, args []string) {
		if buffmanConfigPath != "" {
			handleWithConfig(buffmanConfigPath)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	GenerateCmd.Flags().StringVarP(&buffmanConfigPath, "file", "f", "", "path to buffman config file")
	GenerateCmd.AddCommand(flatbuffersCmd)
}

func handleWithConfig(configPath string) {
	run := runner.NewRunner()
	if err := run.Run(context.Background(), configPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
