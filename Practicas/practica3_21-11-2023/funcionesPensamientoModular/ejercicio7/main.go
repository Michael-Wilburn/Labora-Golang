package main

import "fmt"

func main() {
	invertirString("roma")
}

/*
7.Realice un algoritmo que dado un string lo invierta.
*/

func invertirString(str string) {
	chars := []rune(str)
	var result []rune
	for i := len(chars) - 1; i >= 0; i-- {
		result = append(result, chars[i])
	}
	fmt.Println(string(result))

}
