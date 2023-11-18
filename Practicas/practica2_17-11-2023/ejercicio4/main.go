package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingrese el numero hasta el cual quiere sumar de los primeros numeros naturales")
	fmt.Scan(&num)
	fmt.Printf("La suma de los %v numeros es: %#v\n", num, sumarPrimerosNnaturales(num))
}

//4.Realice un algoritmo para sumar los primeros "n" números naturales, siendo "n" un valor que ingresa el usuario por teclado durante la ejecución del algoritmo

func sumarPrimerosNnaturales(num int) int {
	acumulador := 0
	for i := 1; i < num+1; i++ {
		acumulador += i
	}
	return acumulador
}
