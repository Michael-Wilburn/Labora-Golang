package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(formula(6, 9))
}

/*
Realice un algoritmo que dado dos números calcule el resultado de la potencia del primero elevado al segundo más la sumatoria del primero multiplicado el segundo, todo lo anterior dividido el mínimo común múltiplo entre ambos números.
*/

func maximoComunDivisor(num1, num2 int) int {
	mcd := 0
	afloat := math.Max(float64(num1), float64(num2))
	bfloat := math.Min(float64(num1), float64(num2))
	a := int(afloat)
	b := int(bfloat)

	for {
		if b == 0 {
			break
		}
		mcd = b
		b = a % b
		a = mcd
	}
	return mcd
}

func minimoComunmultiplo(num1, num2 int) int {
	mcm := 0
	afloat := math.Max(float64(num1), float64(num2))
	bfloat := math.Min(float64(num1), float64(num2))
	a := int(afloat)
	b := int(bfloat)

	mcm = (a / maximoComunDivisor(a, b)) * b
	return mcm
}

func potencia(x, y int) int {
	acumulador := 1
	for i := 0; i < y; i++ {
		acumulador *= x
	}
	return acumulador
}

func sumatoria(num int) int {
	if num == 0 {
		return 0
	}
	return num + sumatoria(num-1)
}

func formula(x, y int) int {
	return (potencia(x, y) + (sumatoria(x) * y)) / minimoComunmultiplo(x, y)
}
