package main

import (
	"context"
	"log"

	"github.com/machanirobotics/flatconverter/pkg/converter"
)

func main() {
	protoDir := "deps"
	cleanedDir := "cleaned"
	flatDir := "flatbuffers"
	prefix := "fb"
	c, err := converter.NewConverter(converter.FLATBUFFER, protoDir, cleanedDir, flatDir, prefix)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Convert(context.Background(), true); err != nil {
		log.Fatal(err)
	}
}
