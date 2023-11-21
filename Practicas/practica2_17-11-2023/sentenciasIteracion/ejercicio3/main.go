package main

import "fmt"

func main() {
	fmt.Println(sumarPrimeros3naturales())
	fmt.Println(sumarPrimeros10naturales())

}

func sumarPrimeros3naturales() int {
	acumulador := 0
	for i := 1; i < 4; i++ {
		acumulador += i
	}
	return acumulador
}

func sumarPrimeros10naturales() int {
	acumulador := 0
	for i := 1; i < 11; i++ {
		acumulador += i
	}
	return acumulador

}

//3.Realice un algoritmo para sumar los primeros 3 números naturales. Y luego un algoritmo para sumar los primeros 10 números naturales
