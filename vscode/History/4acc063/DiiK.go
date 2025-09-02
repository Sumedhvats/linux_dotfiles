package golearning

import "math"

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * (0.5)
}
