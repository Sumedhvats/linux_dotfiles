package main

import (
	"net/http"
	"time"
)

func Racer(slowURL string, fastURL string) string {
	startA := time.Now()
	http.Get(slowURL)
	aDuration := time.Since(startA)
	startB := time.Now()
	http.Get(slowURL)
	bDuration := time.Since(startB)
	if aDuration < bDuration {
		return slowURL
	}

	return fastURL
}
