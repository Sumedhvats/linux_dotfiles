package golearning

import "errors"
type Dictionary map[string]string


func (d Dictionary)Search(s string)(string,error){
	value,ok:=d[s]
	if !ok{
		return "",errors.New("")
	}

}