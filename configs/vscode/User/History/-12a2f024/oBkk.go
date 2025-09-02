package golearning
type Dictionary map[string]string


func (d *Dictionary)Search(s string)string{
		d[s]="fadf"
		return d[s]
}