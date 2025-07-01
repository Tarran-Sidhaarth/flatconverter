// package cmd implements the root command for the Buffman CLI application.
// It sets up the main command structure using Cobra, defines global flags,
// and adds subcommands for specific functionalities like 'convert' and 'generate'.
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/machanirobotics/buffman/cmd/completion"
	"github.com/machanirobotics/buffman/cmd/convert"
	"github.com/machanirobotics/buffman/cmd/generate"
	"github.com/machanirobotics/buffman/internal/install"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
// It initializes and runs the tasks defined in the buffman.yml configuration file.
var rootCmd = &cobra.Command{
	Use:   "buffman",
	Short: "A powerful CLI tool for converting and managing buffer schemas.",
	Long: `Buffman is a versatile command-line tool designed to streamline the process
of working with different buffer schemas. It simplifies converting Protocol Buffer
(.proto) files into other formats like FlatBuffers (.fbs), with more formats
planned for the future.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd and
// acts as the primary entry point for the CLI application.
func Execute() {
	// installing flatc if it is missing
	installer := install.NewInstaller(install.FlatbuffersInstaller)
	if !installer.Exists() {
		fmt.Print("flatc is missing. Would you like to install it? (y/n): ")

		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			os.Exit(1)
		}

		// Normalize the response to lowercase for comparison
		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			fmt.Println("Installing flatc...")
			if err := installer.Install(); err != nil {
				fmt.Printf("Failed to install flatc please try installing manually: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("flatc installed successfully!")
		} else {
			fmt.Println("flatc is required for Buffman to work. Exiting...")
			os.Exit(1)
		}
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add subcommands for 'convert' and 'generate' functionalities.
	rootCmd.AddCommand(convert.ConvertCmd, generate.GenerateCmd, completion.CompletionCmd)
}
