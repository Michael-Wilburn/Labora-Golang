package main

import "fmt"

func main() {
	fmt.Println(stringTransform("murcielago"))
}

/*
4.Desarrolle una función que reciba como parámetro una cadena y que reemplace cada vocal por un carácter que la represente simbólicamente. Se puede usar la siguiente tabla de transformación

‘a’ → ‘4’
‘e’ → ‘3’
‘i’ → ‘1’
‘o’ → ‘0’
‘u’ → ‘6’ (no hay mejor candidato, pero si quieren usar otro sean bienvenidos)
Asumiendo que un usuario ingresa “pepa” y quiere transformar su cadena usando la codificación correspondiente, quedaría “p3p4”.
*/

func stringTransform(str string) string {

	longTexto := len(str)
	strToRune := []rune(str)

	for i := 0; i < longTexto; i++ {
		c := strToRune[i]

		switch c {
		case 'a':
			strToRune[i] = '4'
		case 'e':
			strToRune[i] = '3'
		case 'i':
			strToRune[i] = '1'
		case 'o':
			strToRune[i] = '0'
		case 'u':
			strToRune[i] = '6'
		default:
		}
	}
	return string(strToRune)
}
