package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	dynamicstory "github.com/Sumedhvats/Choose-your-own-adventure"
)

func main() {
	file := flag.String("file", "gopher.json", "file name")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *file)
	f, err := os.Open(*file)
	if err != nil {
		panic(err)

	}
	d := json.NewDecoder(f)
	var story dynamicstory.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}
fmt.Printf("%+v\n",story)
}
