package golearning
 Shape Rectangle struct{
	length float64
	breadth float64
 }
 Shape Circle struct{
radius float64
 }
 type Shape interface{
	Ares() float64
 }