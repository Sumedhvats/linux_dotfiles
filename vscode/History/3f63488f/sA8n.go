package main

import (
	"net/http"
	"github.com/Sumedhvats/urlshort/"
	"fmt"

)


func Main(){
mux :=defaultMux()
pathsToUrls:= map[string]string{
	"/github":"https://github.com/sumedhvats/",
	"youtube":"https://www.youtube.com/",
}
mapHandler := urlshort.MapHandler(pathsToUrls,mux)
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}
func defaultMux() *http.ServeMux{
	mux:=http.NewServeMux()
	mux.HandleFunc("/",hello)
	return mux
}