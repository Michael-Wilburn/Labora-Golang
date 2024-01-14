package main

import (
	"errors"
	"fmt"
)

func main() {
	emptyStr := ""
	longStr := "fksjdfhsjdkfhakjfdhjkaldsfhlkasjfhasdkjfhasfkjsadhfhaskjfhsdakfljhsdfkjhsfkjhsdafdaskdgyeuwgrwuyegchbhebcuwbcbcuh"
	fmt.Println(validation(emptyStr))
	fmt.Println(validation(longStr))
}

func validation(str string) error {
	if str == "" {
		return errors.New("\033[31;1m Error: El string esta vacio. \033[0m\n")
	}
	if len(str) > 100 {
		return errors.New("\033[31;1m Error: El string tiene mas de 100 caracateres. \033[0m\n")
	}
	return nil
}

/*
## **Ejercicio 2: Validación de entrada**

Crea una función que acepte un string y devuelva un error si el string está vacío o si es demasiado largo (por ejemplo, más de 100 caracteres).
*/
