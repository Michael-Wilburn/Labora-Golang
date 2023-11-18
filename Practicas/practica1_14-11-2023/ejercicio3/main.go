package main

import (
	"fmt"
)

func main() {
	var temp1 float64
	fmt.Scan(&temp1)
	var temp2 float64
	fmt.Scan(&temp2)
	var temp3 float64
	fmt.Scan(&temp3)

	fmt.Println(ejercicio3(temp1, temp2, temp3))
}

//3.Dados tres valores ambientales de temperatura, desarrollar un algoritmo para calcular e informar la suma y el promedio de dichos valores.

func ejercicio3(temp1 float64, temp2 float64, temp3 float64) (float64, float64) {
	suma := temp1 + temp2 + temp3
	promedio := (temp1 + temp2 + temp3) / 3

	return suma, promedio

}
