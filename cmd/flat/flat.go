package flat

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/machanirobotics/flatconverter/pkg/converter"
	"github.com/machanirobotics/flatconverter/pkg/generator"
	"github.com/spf13/cobra"
)

var FlatCmd = &cobra.Command{
	Use:   "flat",
	Short: "Convert proto files to flatbuffer files",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := converter.NewConverter(converter.FLATBUFFER, protoDir, cleanedDir, flatbufferDir, "")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := c.Convert(context.Background(), true); err != nil {
			fmt.Printf("failed to convert to flat buffer %v", err)
			os.Exit(1)
		}
		packagePrefix := make(map[generator.Languages]string)
		langs := []generator.Languages{}
		for lang := range strings.SplitSeq(strings.TrimSpace(languages), ",") {
			lang, err := generator.StringToLanguage(lang)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			langs = append(langs, lang)
			switch lang {
			case generator.GO:
				packagePrefix[lang] = goModule
			case generator.JAVA:
				packagePrefix[lang] = javaPackage
			default:
				packagePrefix[lang] = ""
			}
		}
		g, err := generator.NewGenerator(generator.FLATBUFFER, flatbufferDir, generatedDir, packagePrefix)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := g.Generate(context.Background(), langs); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
