package golearning
 type Rectangle struct{
	length float64
	breadth float64
 }
 type Circle struct{
radius float64
 }
 type Shape interface{
	Ares() float64
 }