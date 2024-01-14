package main

import "fmt"

func main() {
	fmt.Println(sumarPrimerosImpares(50))
}

func sumarPrimerosImpares(n int) int {
	sum := 0
	odd := 1
	for i := 0; i < n; i++ {
		sum += odd
		odd += 2
	}
	return sum
}

/*
## Ejercicio 1

Escribir un algoritmo para sumar los primeros “n” números impares, siendo un “n” un parámetro en un archivo llamado “odds.go”. Luego se desea poner a prueba la correctitud escribiendo en un archivo “odds_test.go”,se pide usar el enfoque de table driven test. Como sugerencia se puede pensar en que cada test puede describirse como:
*/
