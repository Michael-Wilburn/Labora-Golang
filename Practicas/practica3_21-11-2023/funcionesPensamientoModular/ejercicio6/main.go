package main

import "fmt"

func unAlgoritmo(a int) int {
	quieroCunfundir := a * 2
	for a != 0 {
		a--
		quieroCunfundir += a * 2
	}
	return quieroCunfundir
}

func unaCosa(valor int) {
	resultado := unAlgoritmo(valor) * unAlgoritmo(valor)
	fmt.Println("El resultado es:", resultado)
}

func main() {
	unaCosa(2)
}
