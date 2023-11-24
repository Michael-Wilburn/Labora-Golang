package main

import "fmt"

func main() {
	sonAmigos(220, 284)
}

/*
2. Realice un algoritmo para determinar si dos números son amigos. Un número es amigo de otro cuando la suma de sus divisores propios es igual al otro número.
Sean X1, X2, XN todos divisores propios de X
Sean Y1, Y2, YN todos divisores propios de Y
Si X e Y son amigos entonces X1 + x2 +…+ XN es igual a Y e Y1 + Y2 +…+ XN es igual a X

Ejemplo:
El par (220, 284), ya que:
Los divisores propios de 220 son 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 y 110, que suman 284.
Los divisores propios de 284 son 1, 2, 4, 71 y 142, que suman 220.
*/

func sonAmigos(num1, num2 int) {
	_, acum1 := esPerfecto(num1)
	_, acum2 := esPerfecto(num2)

	if acum1 == num2 && acum2 == num1 {
		fmt.Println("Son amigos")

	} else {
		fmt.Println("No son amigos")
	}

}

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
