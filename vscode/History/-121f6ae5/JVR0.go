package main

import "sync"
type Counter struct{
	mu sync.Mutex
val int;
}
func (c *Counter) Inc(){
	c.val++;
}
func (c *Counter) Value()int{
	return c.val
}