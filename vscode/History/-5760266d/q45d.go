package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	dynamicstory "github.com/Sumedhvats/Choose-your-own-adventure"
)

func main() {
	port := flag.Int("port", 3000, "port to run the http server")
	file := flag.String("file", "gopher.json", "file name")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *file)
	f, err := os.Open(*file)
	if err != nil {
		panic(err)

	}
	story, err := dynamicstory.JsonStory(f)
	if err != nil {
		panic(err)
	}
	h := dynamicstory.NewHandler(story)
	fmt.Println("Starting the server on :", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":% d", *port), h))
}
