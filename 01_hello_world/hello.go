package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const amharic = "Amharic"
const englishHelloPrifix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const amharicHelloPrefix = "ሰላም "

func Hello(name string, language string) string{
	if name == "" {
		name = "World"
	}
	
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case amharic:
		prefix = amharicHelloPrefix
	default:
		prefix = englishHelloPrifix
	}
	return
}

func main() {
	fmt.Println(Hello("", "Spanish"))
}