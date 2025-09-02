package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		return "Hello, World"
	}
	
	if language == "Spanish" {
return "Hola, " + name
	}
	return "Hello"+name
}

func main() {
	fmt.Println(Hello("fdg","English"))
}
