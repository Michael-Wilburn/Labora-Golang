package main

import "fmt"

func main() {
	esPerfecto, acumulador := esPerfecto(6)
	fmt.Printf("El numero esPerfecto: %v\n la suma da como resultado: %v\n", esPerfecto, acumulador)
}

/*1. Realice un algoritmo para determinar si un número es perfecto. Un número es perfecto cuando la suma de sus divisores propios es igual al número. Los divisores propios de un número son todos sus divisores sin contar el mismo número.
Sean X1, X2, XN todos divisores propios de X
Si X es propio entonces X1 + x2 +…+ XN es igual a X
Ejemplo:
6 es un número perfecto
Divisores Propios: 1, 2 y 3.
1 + 2 + 3 = 6
*/

func esPerfecto(num int) (bool, int) {
	var accumulador int
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			accumulador += i
		}
	}
	if num == accumulador {
		return true, accumulador
	}
	return false, accumulador

}
