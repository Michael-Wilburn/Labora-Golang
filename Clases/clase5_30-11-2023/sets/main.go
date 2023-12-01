package main

import (
	"fmt"
)

func main() {
	var ocurrencesByChar map[string]int = make(map[string]int)

	for _, char := range "Hola amigos" {
		_, exists := ocurrencesByChar[string(char)]
		if !exists {
			ocurrencesByChar[string(char)] = 0
		}
		ocurrencesByChar[string(char)]++
	}

	fmt.Println(ocurrencesByChar)

}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
