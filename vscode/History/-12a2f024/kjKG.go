package golearning
type Dictionary map[string]string


func (d *Dictionary)Search(s string)(string,error){
		return (*d)[s],nil
}