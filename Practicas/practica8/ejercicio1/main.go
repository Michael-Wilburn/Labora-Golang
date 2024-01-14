package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerador float64
	var denominador float64
	fmt.Println("Para realizar la división primero ingrese el numerador")
	fmt.Scan(&numerador)
	fmt.Println("Ahora ingrese el denominador")
	fmt.Scan(&denominador)

	resultado, err := division(numerador, denominador)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.2f\n", resultado)
}

func division(numerador, denomiandor float64) (float64, error) {
	if denomiandor != 0 {
		return numerador / denomiandor, nil
	}
	errorCero := errors.New("\033[31;1m La división por cero no es posible \033[0m\n")
	return 0, errorCero
}

/*
## Ejercicio 1: Validación de entrada

Escribe un programa en Go que solicite al usuario dos números (numerador y denominador), intente realizar la división y maneje el caso en el que el denominador sea cero. Si el denominador es cero, imprime un mensaje de error indicando que la división no es posible. En caso contrario, imprime el resultado de la división con dos decimales.

*/
