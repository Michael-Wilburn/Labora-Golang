package main

import (
	"fmt"
	"time"
)

func main() {
	sleepSort()
}

/*
# Ejercicio 1
Usando la técnica del sleep sort, hagamos un programa que imprima los primeros 10 números enteros en orden ascendiente.
¿ Es posible hacerlo en orden descendiente?
*/
func sleepSort() {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
