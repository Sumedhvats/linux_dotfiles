package main

import (
	"testing"
)

func TestCounter(t *testing.T) {
t.Run("Test 3 count increase",func(t *testing.T) {
	counter:=Counter{}
	counter.Inc()
})
	
}
