package main

import (
	"fmt"
	"strconv"
)

func main() {
	// casos capicua
	fmt.Println(capicua(9889))
	fmt.Println(capicua(20102))
	fmt.Println(capicua(335533))

	//casos no capicua
	fmt.Println(capicua(9881))
	fmt.Println(capicua(20103))
	fmt.Println(capicua(335536))
}

/* 5. Realice un algoritmo que dado un número te diga si es capicua.
    a. ¿Hay alguna relación entre estos últimos dos ejercicios?
		a)Si hay una relacion la unica diferencia es que en el palindromo hay que retirar los espacios vacios entre palagras y con los numero capicua no y tener en cuenta que son diferentes tipos de datos con los que se trabja.

*/

func capicua(num int) bool {
	// Convierto el numero a String
	numStr := strconv.Itoa(num)
	//covierto el string a rune
	chars := []rune(numStr)
	var result []rune
	for i := len(chars) - 1; i >= 0; i-- {
		result = append(result, chars[i])
	}

	//devuelve verdadero si es capicua cualquier numero ingresado
	if numStr == string(result) {
		return true
	}

	return false

}
