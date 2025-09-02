package main

import "fmt"

func Hello(name string) string {
	return "Hello, "+name
}
func Hello() string {
	return "Hello, World"
}

func main() {
	fmt.Println(Hello("fdg"))
}