
//package name
package main
import "testing"
import "fmt"
//test adder
func TestAdder(t *testing.T){
	got:=Add(1,2)
	want:=3
	if got!=want{
		t.Errorf("expected %d but got %d",want,got)
	}
}
//example adder
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//  Output: 6
}