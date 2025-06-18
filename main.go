package main

import (
	"context"
	"log"

	"github.com/machanirobotics/flatconverter/pkg/converter"
)

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/bufbuild/protocompile"
// 	"github.com/bufbuild/protocompile/reporter"
// )

func main() {
	c := converter.NewConverter("deps/", "cleaned/")
	err := c.Clean(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	err = c.Convert(context.Background(), "flatbuffers/")
	if err != nil {
		log.Fatal(err)
	}
}
