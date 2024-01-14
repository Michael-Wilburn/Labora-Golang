package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "122211"
	result := func(s string) int {
		num, err := strconv.Atoi(s)

		if err != nil {
			defer func() {
				error := recover()
				fmt.Printf("No se puede tranformar la cadena, el error es --> %v", error)
			}()
			panic(err)
		}
		return num
	}(s)

	fmt.Println(result)

}

/*
## **Ejercicio 3: Conversión de tipos con manejo de errores**

Escribe una función que convierta un string a un número entero (**`int`**). La función debe devolver un error si el string no puede convertirse a un número. Prueba tu función con diferentes strings, incluyendo algunos que no sean números. Usa `panic` y `recover` para manejar los errores
*/
