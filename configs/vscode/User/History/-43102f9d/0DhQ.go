package golearning

 type Shape interface{
	Area() float64
 }
type Rectangle struct {
    Width  float64
    Height float64
}

type Circle struct {
    Radius float64

}
type Triangle struct {
    height float64
    base float64
}
