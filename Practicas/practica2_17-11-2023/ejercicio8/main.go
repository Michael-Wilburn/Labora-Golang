package main

import (
	"fmt"
	"math"
)

func main() {
	//Ej.: mcm (6,9) = 18, mcm (10,15) = 30, mcm (7,14) = 14, mcm (3,7) = 21
	fmt.Println(minimoComunmultiplo(6, 9))
	fmt.Println(minimoComunmultiplo(10, 15))
	fmt.Println(minimoComunmultiplo(3, 7))

}

// 8.Desarrollar un algoritmo para hallar el mínimo común múltiplo (abreviado como mcm) entre dos números naturales. El mínimo común múltiplo entre dos números es el menor número que es divisible por ambos.
func minimoComunmultiplo(num1, num2 int) int {
	mcm := 0
	afloat := math.Max(float64(num1), float64(num2))
	bfloat := math.Min(float64(num1), float64(num2))
	a := int(afloat)
	b := int(bfloat)

	mcm = (a / maximoComunDivisor(a, b)) * b
	return mcm
}

func maximoComunDivisor(num1, num2 int) int {
	mcm := 0
	afloat := math.Max(float64(num1), float64(num2))
	bfloat := math.Min(float64(num1), float64(num2))
	a := int(afloat)
	b := int(bfloat)

	for {
		if b == 0 {
			break
		}
		mcm = b
		b = a % b
		a = mcm
	}
	return mcm
}
