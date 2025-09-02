package main

import (
	"net/http"
	"time"
)

func Racer(slowURL string, fastURL string) string {

	select{
		case <-ping(slowURL):return slowURL
		case <-ping(fastURL):return fastURL
	}
}
// func measureTimeSince(url string) time.Duration {
// 	startA := time.Now()
// 	http.Get(url)
// 	return time.Since(startA)
// }
func ping(url string){
	
}