package main

import "fmt"

func main() {
	fmt.Println("Ingrese los numeros que quiera sumar y obtendra el resultado cuando ingrese 0 ")
	fmt.Println(suma())
}

//5.Redactar un algoritmo que pida al usuario el ingreso de una serie de números reales e imprima por pantalla el resultado de sumarlos. La serie finaliza cuando el usuario ingresa el número cero. Comparando el ejercicio anterior, qué analogías y diferencias encuentra. Debata con sus compañeros, siguiendo las premisas de no violencia. Recuerde, "dime de un hombre sabio y sabré que no es un hombre que recurre a la violencia"

func suma() int {
	var accumulador int
	var num int

	for {
		fmt.Scan(&num)
		if num == 0 {
			return accumulador
		}
		accumulador += num
	}
}
