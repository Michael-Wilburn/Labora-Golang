package main

import "fmt"

func main() {
	esPrimo(3)
	esPrimo(7)
	esPrimo(11)
	esPrimo(4)
}

/*
3.Escriba un algoritmo para determinar si un número es primo. Recordar que número primo es aquel que es divisible por solo por 1 y por si mismo.}
*/

// Si se puede usar la funcion del punto 1 para resolverlo

func esPrimo(num int) {
	_, acumulador := esPerfecto(num)

	if acumulador == 1 {
		fmt.Println("Es un numero Primo")
	} else {
		fmt.Println("No es un numero Primo")
	}
}

func esPerfecto(num int) (bool, int) {
	var accumulador int
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			accumulador += i
		}
	}
	if num == accumulador {
		return true, accumulador
	}
	return false, accumulador

}
