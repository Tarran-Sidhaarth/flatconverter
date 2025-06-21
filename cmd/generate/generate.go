// package generate implements the "generate" command and its subcommands
// for the Buffman CLI. It serves as a container for different code
// generation functionalities.
package generate

import (
	"github.com/spf13/cobra"
)

// GenerateCmd represents the base command for the "generate" functionality.
// It groups all code generation subcommands, such as 'flatbuffers'.
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Contains subcommands for generating code from schema files",
	Long: `The 'generate' command provides access to various code generators.
	
Use a subcommand to specify the type of code to generate. For example, to
generate language-specific files from FlatBuffer schemas, use the 'flatbuffers'
subcommand.`,
	// A Run function is not needed here since this command only serves
	// as a parent for its subcommands.
}

func init() {
	GenerateCmd.AddCommand(flatbuffersCmd)
}
