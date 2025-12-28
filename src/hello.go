package main

import "fmt"

const (
	french  = "french"
	spanish = "spanish"

	english_prefix = "hello"
	spanish_prefix = "hola"
	french_prefix  = "bonjour"
)

func deriveLanguagePrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanish_prefix
	case french:
		prefix = french_prefix
	default:
		prefix = english_prefix
	}

	return
}

func HelloWorld(part string, language string) string {
	if part == "" {
		part = "world"
	}

	return deriveLanguagePrefix(language) + " " + part
}

func main() {
	fmt.Println(HelloWorld("", ""))
}
