package golearning

import (
	
	"testing"
)
	func TestArea(t *testing.T) {
		areaTest:=[]struct{
			shape Shape
			want float64
		}{
			{Rectangle{12,6},72.0},
   			{Circle{10}, 314.1592653589793},

		}

	}