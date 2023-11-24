package main

import "fmt"

// Estructuras de datos

func main() {
	ints := [3]int{1, 2, 3}

	for index, value := range ints {
		fmt.Printf("Indice: %d, Valor: %d\n", index, value)
	}

}
