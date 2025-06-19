package main

import (
	"context"
	"log"

	"github.com/machanirobotics/flatconverter/pkg/converter"
)

func main() {
	protoDir := "/home/tarransidhaarth/Desktop/bruh/temp/test/deps"
	cleanedDir := "/home/tarransidhaarth/Desktop/bruh/temp/test/cleaned"
	flatDir := "/home/tarransidhaarth/Desktop/bruh/temp/test/flatbuffers"
	prefix := ""
	c, err := converter.NewConverter(converter.FLATBUFFER, protoDir, cleanedDir, flatDir, prefix)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Convert(context.Background(), true); err != nil {
		log.Fatal(err)
	}
}
