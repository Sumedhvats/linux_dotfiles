package main

import (
	"net/http"
	"time"
)

func Racer(slowURL string, fastURL string) string {
	startA := time.Now()
	http.Get(slowURL)
	
	aDuration := time.Since(startA)
	println(aDuration)
	startB := time.Now().Minute()
	http.Get(slowURL)
	bDuration := time.Since(startB)
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
