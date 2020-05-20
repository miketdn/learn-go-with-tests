package main

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishPrefix
	case "French":
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	println("Hello, world")
}
