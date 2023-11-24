package main

import "fmt"

func main() {
	fmt.Println(factorial(3))
	fmt.Println(sumatoria(3))
	fmt.Println(multiplicarFactorialPorSumatoria(3))
}

/*
10. Realice un algoritmo para multiplicar el factorial de un número por su sumatoria.

Definimos a la factorial de un número “x” como el resultado de multiplicar sucesivamente los números naturales desde el 1 hasta el número en cuestión.

*/

func multiplicarFactorialPorSumatoria(num int) int {
	return factorial(num) * sumatoria(num)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func sumatoria(num int) int {
	if num == 0 {
		return 0
	}
	return num + sumatoria(num-1)
}
