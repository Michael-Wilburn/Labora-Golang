package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var enteros Numbers
	enteros.Add(1)
	enteros.Add(2)
	enteros.Add(3)
	enteros.RemoveAt(1)
	enteros.Set(0, 0)
	fmt.Println(enteros.ToString())
	fmt.Println(enteros)
	fmt.Println(enteros.Lenght())
}

type Numbers struct {
	conjunto []int
}

func (numbers *Numbers) Add(value int) {
	numbers.conjunto = append(numbers.conjunto, value)
}

func (numbers *Numbers) RemoveAt(index int) bool {
	if index > len(numbers.conjunto) || index < 0 {
		return false
	}
	numbers.conjunto = append(numbers.conjunto[:index], numbers.conjunto[index+1:]...)
	return true
}

func (numbers *Numbers) Lenght() int {
	return len(numbers.conjunto)
}

func (numbers *Numbers) Set(value int, index int) bool {
	if index > len(numbers.conjunto) || index < 0 {
		return false
	}
	numbers.conjunto[index] = value
	return true
}

func (numbers *Numbers) ToString() string {
	valuesText := []string{}
	for i := range numbers.conjunto {
		number := numbers.conjunto[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}
	result := strings.Join(valuesText, ", ")
	return result
}

/*
Defina una estructura de datos para manejar conjuntos de enteros de la siguiente forma.
 `Add(value int) int`
Agrega al final. Se incrementa en uno la longitud.
`RemoveAt(index int) bool`
True si lo borró, false si no pudo (no importa el motivo). Si se borra entonces se decrementa en uno la longitud
`Lenght() int`
Longitud: Cantidad de elementos que tiene. Esta determinada por la cantidad de valores que se agregan menos aquellos que se borran.
`Set(value int, index int) bool`
True si lo modificó, false si no pudo (no importa el motivo)
`ToString() string`
Libre elección. La función strings.Join (del paquete strings) puede ser de útil.
*/
