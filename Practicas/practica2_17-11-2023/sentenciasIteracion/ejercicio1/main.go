package main

import "fmt"

func main() {

	fmt.Printf("El mayor de los 3 numeros es: %#v\n", mayorDe3())
	fmt.Printf("El mayor de los 10 numeros es: %#v\n", mayorDe10())

}

//1.Realizar un algoritmo para determinar el mayor de 3 números. Y luego para determinar el mayor de 10 numeros.

func mayorDe3() int {
	var num int
	maxNum := new(int)

	fmt.Println("Ingrese de a un número y el programa le dira cual es mayor de los tres")
	for i := 0; i < 3; i++ {
		fmt.Scan(&num)
		if num > *maxNum {
			*maxNum = num
		}

	}

	return *maxNum
}

func mayorDe10() int {
	var num int
	maxNum := new(int)

	fmt.Println("Ingrese de a un número y el programa le dira cual es mayor de los diez")
	for i := 0; i < 10; i++ {
		fmt.Scan(&num)
		if num > *maxNum {
			*maxNum = num
		}

	}

	return *maxNum

}
