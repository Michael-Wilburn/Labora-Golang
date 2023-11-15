package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(ejercicio2(num))
}

//2. Escribir un algoritmo que determine si un número es par.

func ejercicio2(n int) string {

	if n%2 == 0 {
		return "Es par"
	}
	return "Es impar"

}
