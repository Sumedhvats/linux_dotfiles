package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Test 3 count increase", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		if(counter.Value()!=3){
			t.Errorf("got %d, want %d", counter.Value(), 3)
		}
	})
t.Run("it runs safely concurrently", func(t *testing.T) {
	wantedCount:=1000
	counter:=Counter{}
	var wg sync.WaitGroup
	wg.Add(wantedCount)
	for i:=0;i<wantedCount;i++{
		go func ()  {
			counter.Inc()
			print(counter.Value())
			wg.Done()
		}()
	}
	wg.Wait()
})
}
