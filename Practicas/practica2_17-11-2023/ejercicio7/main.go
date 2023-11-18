package main

import "fmt"

func main() {
	fmt.Println("Ingrese los dos numeros para encontra el máximo común divisor")
	var num1 int
	var num2 int
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	divisores := divisoresComunes(num1, num2)
	fmt.Printf("Maximo comun divisor es: %v\n", divisores[0])
}

//7.Desarrollar un algoritmo para hallar el máximo común divisor (abreviado como mcd) entre dos números naturales. El máximo común divisor entre dos números es el mayor número que divide a ambos.

func divisoresComunes(num1 int, num2 int) []int {
	var arr1 []int
	var arr2 []int
	var arr3 []int

	for i := num1; i > 0; i-- {
		if num1%i == 0 {
			arr1 = append(arr1, i)
		}
	}

	for i := num2; i > 0; i-- {
		if num2%i == 0 {
			arr2 = append(arr2, i)
		}
	}

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				arr3 = append(arr3, arr1[i])
			}
		}
	}

	return arr3

}
