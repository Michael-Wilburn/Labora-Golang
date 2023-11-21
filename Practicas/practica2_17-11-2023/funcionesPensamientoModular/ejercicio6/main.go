package main

import (
	"fmt"
	"strings"
)

func main() {
	esPalindromo("anita lava la tina")
}

/*
6. Realice un algoritmo que dado un string te diga si es palindromo
*/

func esPalindromo(str string) {
	chars := []rune(str)
	var result []rune
	for i := len(chars) - 1; i >= 0; i-- {
		result = append(result, chars[i])
	}

	reverseStr := strings.ReplaceAll(string(result), " ", "")
	strTrimed := strings.ReplaceAll(str, " ", "")

	if strTrimed == reverseStr {
		fmt.Println("La cadena de caracteres es un palindromo")
	} else {
		fmt.Println("La cadena de caracteres no es un palindromo")
	}

}
