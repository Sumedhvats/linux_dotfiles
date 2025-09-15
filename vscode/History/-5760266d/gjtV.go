package main

import (
	"flag"
	"fmt"
)

func main() {
	file := flag.String("file", "gopher.json", "file name")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *file)
}
