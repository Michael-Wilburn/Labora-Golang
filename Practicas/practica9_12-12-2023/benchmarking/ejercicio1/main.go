package main

func main() {

}

func sumaOddsOne(n int) int {
	sum := 0
	odd := 1
	for i := 0; i < n; i++ {
		sum += odd
		odd += 2
	}
	return sum
}

func sumaOddsTwo(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += 2*i - 1
	}
	return sum
}

func sumaOddsThree(n int) int {
	return n * n
}
func sumaOddsFour(n int) int {
	return n*(n+1) - n
}

/*
## Ejercicio 1

Se pide implementar al menos dos formas de sumar los primeros “n” números impares (una de ellas deberías tener por el ejercicio 1 de esta guía). Y luego comparar cual es más eficiente.

*/
