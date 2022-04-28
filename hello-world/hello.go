package main

import "fmt"

// private constant
const helloMessageEnglishPrefix = "Hello, "
const helloMessageBrazilianPortuguesePrefix = "E aí, "
const helloMessageYorubaPrefix = "Pẹlẹ o, "
const helloMessageDefaultEnglishSuffix = "World"
const helloMessageDefaultBrazilianPortugueseSuffix = "man"
const helloMessageDefaultYorubaSuffix = "Aye"

// function with named return
func generateHelloMessagePrefix(language string) (messagePrefix string) {
	switch language {
	case "Brazilian Portuguese":
		messagePrefix = helloMessageBrazilianPortuguesePrefix
	case "Yoruba":
		messagePrefix = helloMessageYorubaPrefix
	case "English":
		messagePrefix = helloMessageEnglishPrefix
	default:
		messagePrefix = helloMessageEnglishPrefix
	}

	return
}

func generateHelloMessageDefaultSuffix(name, language string) (messageSuffix string) {
	if name == "" {
		switch language {
		case "Brazilian Portuguese":
			messageSuffix = helloMessageDefaultBrazilianPortugueseSuffix
		case "Yoruba":
			messageSuffix = helloMessageDefaultYorubaSuffix
		case "English": messageSuffix = helloMessageDefaultEnglishSuffix
		default:
			messageSuffix = helloMessageDefaultEnglishSuffix
		}
		return
	}

	messageSuffix = name

	return messageSuffix
}

func Hello(name, language string) string {
	return generateHelloMessagePrefix(language) + generateHelloMessageDefaultSuffix(name, language)
}

// entry point for the package
func main() {
	fmt.Println(Hello("world", ""))
}
