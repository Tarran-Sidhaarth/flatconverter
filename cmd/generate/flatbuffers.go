package generate

import (
	"context"
	"fmt"
	"os"

	"github.com/machanirobotics/buffman/internal/configuration"
	"github.com/machanirobotics/buffman/internal/generate"
	"github.com/machanirobotics/buffman/internal/generate/language"
	"github.com/machanirobotics/buffman/internal/options"
	"github.com/machanirobotics/buffman/internal/runner"
	"github.com/spf13/cobra"
)

// Flags for the 'generate flatbuffers' command.
var (
	// flatbuffersDir specifies the input directory containing .fbs schema files.
	flatbuffersDir string
	// targetDir specifies the output directory for the generated code.
	targetDir string
	// language defines the target programming language for code generation.
	lang string
	// moduleOptions provides language-specific options, like a package prefix.
	moduleOptions string
	//
	buffmanConfigPath string
)

// flatbuffersCmd represents the command to generate code from FlatBuffer schemas.
var flatbuffersCmd = &cobra.Command{
	Use:   "flatbuffers",
	Short: "Generates language-specific source code from FlatBuffer schema files (.fbs)",
	Long: `The 'flatbuffers' command invokes the FlatBuffers compiler (flatc) to generate
source code for one or more target languages from your .fbs schema files.

You must specify the directory containing your .fbs files and the target language.
An output directory can also be specified, otherwise the files will be generated
in the current directory.

Example:
  buffman generate flatbuffers \
    --flatbuffers_dir=./path/to/fbs \
    --target_dir=./gen/go \
    --language=go \
    --module_options="github.com/your-org/your-project/gen/go"`,
	Run: func(cmd *cobra.Command, args []string) {
		g, err := generate.NewGenerate(generate.Flatbuffers)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if buffmanConfigPath != "" {
			run := runner.NewRunner()
			if err := run.Run(context.Background(), buffmanConfigPath); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			if lang == "" || targetDir == "" {
				fmt.Println("please enter a language and output directory or specify a buffman config file")
			} else {
				if err := g.Generate(context.Background(), options.GenerateOptions{
					InputDir: flatbuffersDir,
					LanguagDetails: map[language.Language]options.LanguageGenerateOptions{
						language.Language(lang): {
							OutputDir: targetDir,
							Opts:      []string{parseModuleOptions(language.Language(lang), moduleOptions)},
						},
					},
				}); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Printf("Successfully generated %s files in %s\n", lang, targetDir)
			}
		}
	},
}

func init() {
	// Define flags for the 'generate flatbuffers' command.
	flatbuffersCmd.Flags().StringVarP(&flatbuffersDir, "flatbuffers_dir", "I", "", "Directory containing the source .fbs schema files")
	flatbuffersCmd.Flags().StringVarP(&targetDir, "target_dir", "o", "./", "Output directory for the generated source code")
	flatbuffersCmd.Flags().StringVarP(&lang, "language", "l", "", "Target language for code generation (e.g., go, java, cpp, kotlin)")
	flatbuffersCmd.Flags().StringVarP(&moduleOptions, "module_options", "m", "", "Language-specific options (e.g., Go package path or Java package name)")
	flatbuffersCmd.Flags().StringVarP(&buffmanConfigPath, "file", "f", "", "path to buffman config file")
}

func parseModuleOptions(languageType language.Language, opt string) string {
	commandOpts := configuration.CommandOptionsMap[languageType]
	return commandOpts["flatbuffer"].Options["package"].Flag + " " + opt
}
