package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)
func TestRacer(t *testing.T) {

	slowURL := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(20 * time.Millisecond)
				w.WriteHeader(http.StatusOK)

	}))
	fastURL := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)

	}))

	want := fastURL.URL
	got := Racer(slowURL.URL, fastURL.URL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	slowURL.Close()
	fastURL.Close()
}
func delayedServer(time.Duration) httptest.Server{
	return  *httptest.NewServer(http.HandlerFunc(w http.ResponseWriter,r *http.Request))

}