package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(ejercicio1(num))
}

//1. Realizar un algoritmo para leer un número e informar si es mayor, igual o menos a cero.

func ejercicio1(n int) string {

	switch {
	case n > 0:
		return "El número ingresado es mayor a cero"
	case n < 0:
		return "El número ingresado es menor a cero"
	case n == 0:
		return "El número ingresado es igual a cero"
	default:
		return "Algo salio mal"
	}

}
