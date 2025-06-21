package generate

import (
	"context"
	"fmt"
	"os"

	"github.com/machanirobotics/buffman/pkg/generator"
	"github.com/spf13/cobra"
)

// Flags for the 'generate flatbuffers' command.
var (
	// flatbuffersDir specifies the input directory containing .fbs schema files.
	flatbuffersDir string
	// targetDir specifies the output directory for the generated code.
	targetDir string
	// language defines the target programming language for code generation.
	language string
	// moduleOptions provides language-specific options, like a package prefix.
	moduleOptions string
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
		lang, err := generator.StringToLanguage(language)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		g, err := generator.NewGenerator(generator.FLATBUFFER, flatbuffersDir, targetDir, map[generator.Languages]string{lang: moduleOptions})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := g.Generate(context.Background(), []generator.Languages{lang}); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Successfully generated %s files in %s\n", language, targetDir)
	},
}

func init() {
	// Attach the flatbuffersCmd to its parent, GenerateCmd.
	GenerateCmd.AddCommand(flatbuffersCmd)

	// Define flags for the 'generate flatbuffers' command.
	flatbuffersCmd.Flags().StringVarP(&flatbuffersDir, "flatbuffers_dir", "I", "", "Directory containing the source .fbs schema files")
	flatbuffersCmd.Flags().StringVarP(&targetDir, "target_dir", "o", "./", "Output directory for the generated source code")
	flatbuffersCmd.Flags().StringVarP(&language, "language", "l", "", "Target language for code generation (e.g., go, java, cpp, kotlin)")
	flatbuffersCmd.Flags().StringVarP(&moduleOptions, "module_options", "m", "", "Language-specific options (e.g., Go package path or Java package name)")

	// Mark essential flags as required.
	flatbuffersCmd.MarkFlagRequired("flatbuffers_dir")
	flatbuffersCmd.MarkFlagRequired("language")
}
