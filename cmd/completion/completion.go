// this file is overriding the default completion
// of cobra to avoid it being seen in the subcommands
package completion

import "github.com/spf13/cobra"

var CompletionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate the autocompletion script for the specified shell",
}

func init() {
	CompletionCmd.Hidden = true
}
