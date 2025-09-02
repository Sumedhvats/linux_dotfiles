package golearning
import "testing"
func TestSum(t *testing.T){
numbers:=[...]int{1,2,3,4,5}
got := Sum(numbers[:])
want := 12
if got!=want {
	t.Errorf("got %d instaed of %d",got,want)
}
}