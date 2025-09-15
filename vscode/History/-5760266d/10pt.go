package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("file", "gopher.json", "file name")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *file)
	f,err:=os.Open(*file)
	if err!=nil{
		panic(err)

	}
		d:=json.NewDecoder(f)
		var story cyoa.Story
		if err :=d.Decode();err!=nil{

		} 

}
