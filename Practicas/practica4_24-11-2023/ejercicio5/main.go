package main

import (
	"fmt"
)

func main() {

	//Test casos Validos
	fmt.Println(dnaSequence("ATGCGTAT", "ATGCGTAT"))
	fmt.Println(dnaSequence("ATGCGTAT", "TGCGTATA"))
	fmt.Println(dnaSequence("ATGCGTAT", "GCGTATAT"))
	fmt.Println(dnaSequence("ATGCGTAT", "CGTATATG"))

	fmt.Println(dnaSequence("ATGCGTAT", "GTATATGC"))
	fmt.Println(dnaSequence("ATGCGTAT", "TATATGCG"))
	fmt.Println(dnaSequence("ATGCGTAT", "ATATGCGT"))
	fmt.Println(dnaSequence("ATGCGTAT", "TATGCGTA"))

	fmt.Println("---------------------------------")
	//Test casos no validos
	fmt.Println(dnaSequence("ATGCGTAT", "TTGCGTAT"))
	fmt.Println(dnaSequence("ATGCGTAT", "TGCGTACA"))
	fmt.Println(dnaSequence("ATGCGTAT", "GCATATAT"))
	fmt.Println(dnaSequence("ATGCGTAT", "CGTATAAG"))

	fmt.Println(dnaSequence("ATGCGTAT", "ATATATGC"))
	fmt.Println(dnaSequence("ATGCGTAT", "TATATGTC"))
	fmt.Println(dnaSequence("ATGCGTAT", "ATATGCCT"))
	fmt.Println(dnaSequence("ATGCGTAT", "TCTGCGTA"))

}

/*
Una cadena de ADN se representa como una secuencia circular de bases (adenina, timina, citosina y guanina) que es única para cada ser vivo, por ejemplo:
1.
| A | T | G |
| T |   | C |
| A | T | G |
Dicha cadena se puede representar como un arreglo de caracteres recorriendola en sentido horario desde la parte superior izquierda:
A T G C G T A T
Se pide diseñar una clase que represente una secuencia de ADN e incluya un método booleano que nos devuelva **true** si dos cadenas de ADN coinciden.
MUY IMPORTANTE: La secuencia de ADN es cíclica, por lo que puede comenzar en cualquier posición. Por ejemplo, las dos secuencias siguientes coinciden:

**A** T G C G T A T

ATGCGTAT
TGCGTATA


A T **A** T G C G T

Pregunta existencial: ¿la cantidad de combinaciones de ADN debe ser finita? ¿Cuántas combinaciones diferentes puede haber? ¿Puede ocurrir algún día que se empiezan a repetir??

ATATGCGT
*/

func dnaSequence(sequence1, sequence2 string) bool {

	seqStr1 := []rune(sequence1)
	seqStr2 := []rune(sequence2)

	for i := 0; i < len(sequence1); i++ {
		acumulador := 0
		if seqStr1[0] == seqStr2[0] {
			for j := 0; j < len(sequence1); j++ {
				if seqStr1[j] == seqStr2[j] {
					acumulador += 1
				}
				if acumulador == len(seqStr1) {
					return true
				}
			}
		}

		first := seqStr1[0]
		seqStr1 = append(seqStr1[:0], seqStr1[1:]...)
		seqStr1 = append(seqStr1, first)
	}

	return false
}
