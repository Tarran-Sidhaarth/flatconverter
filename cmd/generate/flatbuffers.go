package generate

import (
	"context"
	"fmt"
	"os"

	"github.com/machanirobotics/buffman/pkg/generator"
	"github.com/spf13/cobra"
)

var (
	flatbuffersDir string
	targetDir      string
	language       string
	moduleOptions  string
)

var flatbuffersCmd = &cobra.Command{
	Use:   "flatbuffer",
	Short: "command used to generate language specific files from the fbs files",
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
	},
}

func init() {
	flatbuffersCmd.Flags().StringVarP(&flatbuffersDir, "flatbuffers_dir", "I", "", "Directory where the fbs files are stores")
	flatbuffersCmd.Flags().StringVarP(&targetDir, "target_dir", "o", "./", "Directory where the generated files need to be placed")
	flatbuffersCmd.Flags().StringVarP(&language, "language", "l", "", "Language (go, java, cpp, kotlin)")
	flatbuffersCmd.Flags().StringVarP(&moduleOptions, "module_options", "m", "", "Prefix of the package or module")
	flatbuffersCmd.MarkFlagRequired("flatbuffers_dir")
	flatbuffersCmd.MarkFlagRequired("language")
}
