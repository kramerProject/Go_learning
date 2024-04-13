package main

import "fmt"

const french = "french"
const spanish = "spanish"
const prefixHelloPortuguese = "Olá, "
const prefixHelloSpanish = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(name, language string) string {
	if name == "" {
		name = "Mundo"
	}
	return helloPrefix(language) + name + "."
}

func helloPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = prefixHelloSpanish
	case french:
		prefix = prefixoOlaFrances
	default:
		prefix = prefixHelloPortuguese
	}
	return
}
func main() {
	fmt.Println(Ola("Kramer", "spanish"))
}
