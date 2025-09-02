package main

import (
	"net/http"
	"time"
)
func Racer(slowURL string,fastURL string)string{
	startA:=time.Now();
	http.Get(slowURL);
	aDuration:=time.Now()-startA;
}