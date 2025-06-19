package flat

var (
	languages     string
	protoDir      string
	cleanedDir    string
	flatbufferDir string
	generatedDir  string
	goModule      string
	javaPackage   string
)

func init() {
	FlatCmd.Flags().StringVarP(&languages, "language", "l", "", "supported languages(go, cpp, java, kotlin)")
	FlatCmd.Flags().StringVarP(&protoDir, "proto-dir", "p", "", "Directory containing proto files")
	FlatCmd.Flags().StringVarP(&cleanedDir, "cleaned-dir", "c", "", "Directory for cleaned proto files")
	FlatCmd.Flags().StringVarP(&flatbufferDir, "flatbuffer-dir", "f", "", "Directory for flatbuffer output")
	FlatCmd.Flags().StringVarP(&generatedDir, "generated-dir", "g", "", "Directory for generated files")
	FlatCmd.Flags().StringVar(&goModule, "go-module", "", "The go module prefix to append to the files")
	FlatCmd.Flags().StringVar(&javaPackage, "java-package-prefix", "", "Prefix to attach to the java files")
	FlatCmd.MarkFlagRequired("language")
	FlatCmd.MarkFlagRequired("proto-dir")
}
