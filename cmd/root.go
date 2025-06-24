// package cmd implements the root command for the Buffman CLI application.
// It sets up the main command structure using Cobra, defines global flags,
// and adds subcommands for specific functionalities like 'convert' and 'generate'.
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/machanirobotics/buffman/cmd/convert"
	"github.com/machanirobotics/buffman/cmd/generate"
	"github.com/machanirobotics/buffman/internal/install"
	"github.com/machanirobotics/buffman/pkg/runner"
	"github.com/spf13/cobra"
)

// configPath stores the path to the configuration file, which can be set via a command-line flag.
var configPath string

// rootCmd represents the base command when called without any subcommands.
// It initializes and runs the tasks defined in the buffman.yml configuration file.
var rootCmd = &cobra.Command{
	Use:   "buffman",
	Short: "A powerful CLI tool for converting and managing buffer schemas.",
	Long: `Buffman is a versatile command-line tool designed to streamline the process
of working with different buffer schemas. It simplifies converting Protocol Buffer
(.proto) files into other formats like FlatBuffers (.fbs), with more formats
planned for the future.

All operations are driven by a central configuration file, typically 'buffman.yml',
which defines the conversion and generation tasks. If no subcommand is specified,
Buffman will run all tasks defined in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := runner.NewRunner()
		if err := r.Run(context.Background(), configPath); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd and
// acts as the primary entry point for the CLI application.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
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
	// Define a persistent flag for the configuration file path.
	// This flag is available to the root command and all its subcommands.
	rootCmd.Flags().StringVarP(&configPath, "file", "f", "buffman.yml", "Path to the Buffman configuration file")

	// Add subcommands for 'convert' and 'generate' functionalities.
	rootCmd.AddCommand(convert.ConvertCmd, generate.GenerateCmd)
}
