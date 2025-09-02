package golearning
 type Rectangle Shape{
	length float64
	breadth float64
 }
 type Circle Shape{
radius float64
 }
 type Shape interface{
	Ares() float64
 }