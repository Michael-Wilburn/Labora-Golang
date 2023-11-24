package main

import (
	"fmt"
)

func toInt(cad string) int {
	long := len(cad)
	i, pot, res := 0, 1, 0

	for i < long {
		c := cad[long-i-1]
		res += int(c-'0') * pot
		pot *= 10
		i++
	}

	return res
}

func main() {
	var cadena string
	fmt.Print("Ingrese un número como cadena: ")
	fmt.Scan(&cadena)

	resultado := toInt(cadena)
	fmt.Printf("El resultado como entero es: %d\n", resultado)
}
