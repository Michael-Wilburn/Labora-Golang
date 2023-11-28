package main

import "fmt"

func contarLetra(letra rune, texto string) int {
	cant := 0
	longTexto := len(texto)

	for i := 0; i < longTexto; i++ {
		c := rune(texto[i])
		if c == letra {
			cant++
		}
	}

	return cant
}

func contarVocales(text string) {
	var cantA, cantE, cantI, cantO, cantU int

	cantA = contarLetra('a', text)
	cantE = contarLetra('e', text)
	cantI = contarLetra('i', text)
	cantO = contarLetra('o', text)
	cantU = contarLetra('u', text)

	fmt.Printf("Repeticiones de a,e,i,o,u son: %d, %d, %d, %d, %d\n", cantA, cantE, cantI, cantO, cantU)
}

func main() {
	contarVocales("algo lindo")
}
