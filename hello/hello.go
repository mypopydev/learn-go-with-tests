package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello return a strings
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}
