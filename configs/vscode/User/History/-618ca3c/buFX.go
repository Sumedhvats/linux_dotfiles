package main

import (
	"net/http"
)

func Racer(slowURL string, fastURL string) string {

	select{
		case <-ping(slowURL):return slowURL
		case <-ping(fastURL):return fastURL
	}
}
func ping(url string)chan struct{}{
	ch :=make(chan struct{})
	go func(){
		http.Get(url)
		close(ch)																																													
	}()
	return ch
}