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

// 9.Dado un número de 5 cifras, determinar si es capicúa. Si fuera un número de 6 cifras ¿Sirve la resolución planteada? ¿Cómo habría que modificarla?
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

// El ejercicio resuelto de esta manera no importa de cuantas cifras sea el numero va a poder determinar si es capicua o no
