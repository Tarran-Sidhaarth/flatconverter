package generate

import (
	"context"
	"fmt"
	"os"

	"github.com/machanirobotics/buffman/internal/configuration"
	"github.com/machanirobotics/buffman/internal/generate"
	"github.com/machanirobotics/buffman/internal/generate/language"
	"github.com/machanirobotics/buffman/internal/options"
	"github.com/spf13/cobra"
)

// Flags for the 'generate flatbuffers' command.
var (
	flatbuffersDir string
	targetDir      string
	lang           string
	moduleOptions  string
)

var flatbuffersCmd = &cobra.Command{
	Use:   "flatbuffers",
	Short: "Generates language-specific source code from FlatBuffer schema files (.fbs)",
	Long: `The 'flatbuffers' command invokes the FlatBuffers compiler (flatc) to generate
source code for one or more target languages from your .fbs schema files.

You must specify the directory containing your .fbs files, target language, and output directory
using the command-line flags.

Examples:
  # Generate Go code from FlatBuffer schemas
  buffman generate flatbuffers \
    --flatbuffers_dir=./path/to/fbs \
    --target_dir=./gen/go \
    --language=go \
    --module_options="github.com/your-org/your-project/gen/go"

  # Generate C++ code from FlatBuffer schemas
  buffman generate flatbuffers \
    --flatbuffers_dir=./schemas \
    --target_dir=./generated/cpp \
    --language=cpp`,
	Run: func(cmd *cobra.Command, args []string) {
		if lang == "" || targetDir == "" {
			fmt.Println("please specify both --language and --target_dir flags")
			cmd.Help()
			return
		}
		handleGenerate(lang, targetDir, flatbuffersDir, moduleOptions)
	},
}

func init() {
	flatbuffersCmd.Flags().StringVarP(&flatbuffersDir, "flatbuffers_dir", "I", "", "Directory containing the source .fbs schema files")
	flatbuffersCmd.Flags().StringVarP(&targetDir, "target_dir", "o", "", "Output directory for the generated source code")
	flatbuffersCmd.Flags().StringVarP(&lang, "language", "l", "", "Target language for code generation (e.g., go, java, cpp, kotlin)")
	flatbuffersCmd.Flags().StringVarP(&moduleOptions, "module_options", "m", "", "Language-specific options (e.g., Go package path or Java package name)")

	// Mark required flags
	flatbuffersCmd.MarkFlagRequired("language")
	flatbuffersCmd.MarkFlagRequired("target_dir")
}

func handleGenerate(lang, targetDir, flatbuffersDir, moduleOptions string) {
	g, err := generate.NewGenerate(generate.Flatbuffers)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = g.Generate(context.Background(), options.GenerateOptions{
		InputDir: flatbuffersDir,
		LanguagDetails: map[language.Language]options.LanguageGenerateOptions{
			language.Language(lang): {
				OutputDir: targetDir,
				Opts:      []string{parseModuleOptions(language.Language(lang), moduleOptions)},
			},
		},
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %s files in %s\n", lang, targetDir)
}

func parseModuleOptions(languageType language.Language, opt string) string {
	if opt == "" {
		return ""
	}

	commandOpts, ok := configuration.CommandOptionsMap[languageType]
	if !ok {
		return ""
	}
	fmt.Println("Command Options: ", commandOpts["flatbuffer"].Options["package"].Flag)
	return commandOpts["flatbuffer"].Options["package"].Flag + " " + opt
}
