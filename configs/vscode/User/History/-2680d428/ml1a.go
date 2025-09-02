package golearning

import (
	"testing"
)

func TestArea(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}
	for _, tt := range areaTest {
		got:=tt.shape.Area()
		if got!=tt.want{
			t.Errorf("got %f instead of %f",got,tt.want)
		}
	}
}
