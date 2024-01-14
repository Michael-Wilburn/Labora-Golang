package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := rand.Perm(100000000)
	fmt.Println(sumaSecuencial(s))
	fmt.Println(sumaConcurrente(s))
}

func sumaSecuencial(n []int) int {
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return sum
}
func sumaConcurrente(n []int) int {
	mitadUno := n[:(len(n) / 2)]
	mitadDos := n[(len(n) / 2):]

	canal := make(chan int)

	go func() {
		canal <- sumaSecuencial(mitadUno)
	}()
	go func() {
		canal <- sumaSecuencial(mitadDos)
	}()

	sum1 := <-canal
	sum2 := <-canal

	close(canal)
	return sum1 + sum2
}

/*
## Ejercicio 2
Hemos usado canales para sumar de forma concurrente todos los valores presentes en los elementos de un slice (o array) de números enteros. ¿Qué te parece si ponemos a prueba si la concurrencia en este caso es realmente efectiva? Bueno, es una pregunta pero lo tenés que hacer igual… lo lamento 😁 y me gustaría que discutamos luego sobre CUANDO la concurrencia mejora la performance. Recordar que concurrencia no es lo mismo que paralelismo!.
*/
