package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

var helloPrefix string

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		helloPrefix = spanishHelloPrefix
	} else {
		helloPrefix = englishHelloPrefix
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
