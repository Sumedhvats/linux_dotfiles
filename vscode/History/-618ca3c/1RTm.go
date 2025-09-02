package main

import (
	"net/http"
	"time"
)

func Racer(slowURL string, fastURL string) string {
	startA := time.Now()
	http.Get(slowURL)
	
	aDuration := time.Since(startA).Milliseconds()
	println(aDuration)
	startB := time.Now()
	http.Get(slowURL)
	bDuration := time.Since(startB).Milliseconds()
		println(bDuration)
	if aDuration < bDuration {
		return slowURL
	}

	return fastURL
}
func main(){
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"
	print(Racer(slowURL,fastURL))
}
