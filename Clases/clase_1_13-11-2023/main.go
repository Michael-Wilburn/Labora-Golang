package main

import "fmt"

var GLOBAL int = 2

func main() {
	var THE_BEATLES int = 2
	THE_ROLLING_STONES := 3
	var ALGO bool = true
	var caracter rune = 'a'
	var oracion string = "Una frase larga..."
	const pepe int = 2

	fmt.Println(THE_BEATLES, THE_ROLLING_STONES, ALGO, caracter, oracion)

}
