package main

import (
	"net/http"
	"time"
)

func Racer(slowURL string, fastURL string) string {
	
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
func measureTimeSince(url string) time.Duration{
startA := time.Now()
	http.Get(url)
	
time.Since(startA)
	return 
}
