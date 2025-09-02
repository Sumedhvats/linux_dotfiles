package golearning

 type Shape interface{
	Area() float64
 }
  type Rectangle Shape{
	length float64
	breadth float64
 }
 type Circle Shape{
radius float64
 }