package main

import "net/http"


func Main(){
mux :=defaultMux()
pathsToUrls:= map[string]string{
	"/github":"https://github.com/sumedhvats/",
	"youtube":"https://www.youtube.com/",
}
mapHandler := urlShort.MapHandler(pathsToUrls,mux)
}
func defaultMux() *http.ServeMux{
	mux:=http.NewServeMux()
	mux.HandleFunc("/",hello)
	return mux
}