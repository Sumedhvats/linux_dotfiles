package main

import (
	"net/http"
	"time"
)

func Racer(slowURL string, fastURL string) string {

	aDuration := measureTimeSince(slowURL)
	bDuration := measureTimeSince(fastURL)
	println(bDuration)
	if aDuration < bDuration {
		return slowURL
	}
	return fastURL
}
func measureTimeSince(url string) time.Duration {
	startA := time.Now()
	http.Get(url)
	return time.Since(startA)
}
