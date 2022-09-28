package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPreix = "Bonjour, "

func main() {
	fmt.Println(Hello("Lucas",""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish"{
		return spanishHelloPrefix + name
	}

	if language == "French"{
		return frenchHelloPreix + name
	}

	return englishHelloPrefix + name
}