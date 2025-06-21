package convert

import (
	"context"
	"fmt"
	"os"

	"github.com/machanirobotics/buffman/pkg/converter"
	"github.com/spf13/cobra"
)

var (
	protoDir      string
	flatbufferDir string
)

var flatbuffersCmd = &cobra.Command{
	Use:   "flatbuffers",
	Short: "command used to convert the proto files to fb files",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := converter.NewConverter(converter.FLATBUFFER, protoDir, flatbufferDir, "")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := c.Convert(context.Background()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	flatbuffersCmd.Flags().StringVarP(&protoDir, "proto_dir", "I", "", "Directory where the protofiles are present")
	flatbuffersCmd.Flags().StringVarP(&flatbufferDir, "output_dir", "o", "./", "Directory where the flatbuffers files need to be placed")
	flatbuffersCmd.MarkFlagRequired("proto_dir")
}
