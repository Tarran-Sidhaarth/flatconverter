package convert

import (
	"context"
	"fmt"
	"os"

	"github.com/machanirobotics/buffman/pkg/converter"
	"github.com/spf13/cobra"
)

// Flags for the 'convert flatbuffers' command.
var (
	// protoDir specifies the input directory containing .proto files.
	protoDir string
	// flatbufferDir specifies the output directory for the generated FlatBuffer files.
	flatbufferDir string
)

// flatbuffersCmd represents the command to convert Protocol Buffer (.proto) files to FlatBuffer (.fbs) files[1].
var flatbuffersCmd = &cobra.Command{
	Use:   "flatbuffers",
	Short: "Converts Protocol Buffer (.proto) files to FlatBuffer (.fbs) schema files",
	Long: `The 'flatbuffers' command converts Protocol Buffer schema files (.proto) into
FlatBuffer schema files (.fbs). This is useful for migrating or enabling
interoperability between these two serialization formats[1].

You must specify the directory containing your source .proto files. The output
directory for the .fbs files can also be specified; otherwise, they are placed
in the current directory by default[1].

Example:
  buffman convert flatbuffers \
    --proto_dir=./path/to/protos \
    --output_dir=./gen/fbs`,
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
		fmt.Printf("Successfully converted .proto files in %s to .fbs files in %s\n", protoDir, flatbufferDir)
	},
}

func init() {
	// Define and attach flags to the 'convert flatbuffers' command[1].
	flatbuffersCmd.Flags().StringVarP(&protoDir, "proto_dir", "I", "", "Directory containing the source .proto files")
	flatbuffersCmd.Flags().StringVarP(&flatbufferDir, "output_dir", "o", "./", "Output directory for the generated FlatBuffer (.fbs) files")

	// Mark the proto_dir flag as required since it's essential for the command to run[1].
	flatbuffersCmd.MarkFlagRequired("proto_dir")
}
