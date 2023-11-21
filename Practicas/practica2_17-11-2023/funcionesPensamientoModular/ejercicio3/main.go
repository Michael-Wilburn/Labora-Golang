package main

import "fmt"

func main() {
	esPrimo(3)
	esPrimo(7)
	esPrimo(11)
	esPrimo(4)
}

/*
Escriba un algoritmo para determinar si un número es primo. Recordar que número primo es aquel que es divisible por solo por 1 y por si mismo.}
*/

func esPrimo(num int) {
	var acumulador int
	for i := 1; i < num; i++ {
		if num%i == 0 {
			acumulador += i
		}
	}

	if acumulador == 1 {
		fmt.Println("Es un numero Primo")
	} else {
		fmt.Println("No es un numero Primo")
	}
}
