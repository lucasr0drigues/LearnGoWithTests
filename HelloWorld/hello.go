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

	prefix := englishHelloPrefix

	switch language{
	case french:
		prefix = frenchHelloPreix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}