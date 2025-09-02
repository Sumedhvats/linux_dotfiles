package golearning

import "math"

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.breadth + rect.length)
}

func (rect Rectangle) Area() float64 {
	return rect.breadth * rect.length
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return t.base * t.height * (0.5)
}
