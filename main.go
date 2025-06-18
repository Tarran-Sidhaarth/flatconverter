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

	c, err := converter.NewConverter(converter.CONVERTER_TYPE_FLATBUFFER, protoDir, cleanedDir, flatDir)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Convert(context.Background(), true); err != nil {
		log.Fatal(err)
	}
}
