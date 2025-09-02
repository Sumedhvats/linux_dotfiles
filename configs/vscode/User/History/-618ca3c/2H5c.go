package main

import (
	"net/http"
)

func Racer(slowURL string, fastURL string) (string,error) {

	select{
		case <-ping(slowURL):return slowURL,nil
		case <-ping(fastURL):return fastURL,_
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