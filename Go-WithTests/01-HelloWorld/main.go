package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const deutschHelloPrefix = "Halo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case "Spanish":
		return spanishHelloPrefix + name
	case "French":
		return frenchHelloPrefix + name
	case "Deutsch":
		return deutschHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
