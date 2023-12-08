// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"math/rand"
)

type IntSequence interface {
	Next() int
}

type LinealSequence struct {
	current int
}

func (s *LinealSequence) Next() int {
	var next int = s.current
	s.current++
	return next
}

type PrimesSequence struct {
	lastPrime int
}

func (s *PrimesSequence) Next() int {
	var next int
	for n := s.lastPrime + 1; n <= math.MaxInt32; n++ {
		if isPrime(n) {
			next = n
			break
		}
	}
	s.lastPrime = next
	return next
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var intSequence IntSequence

	if rand.Intn(2) == 0 {
		intSequence = &LinealSequence{} // aca tienen que asignar la dirección pues el método Next() opera sobre *Sequence y no sobre Sequence (recordar que son dos tipos de datos distintos)
	} else {
		intSequence = &PrimesSequence{} // idem, hay que asignar la dirección pues trabajará un puntero a PrimesSequence
	}
	for i := 1; i <= 50; i++ {
		fmt.Println(intSequence.Next())
	}

}

/*
(1) Se necesita generar dos secuencias de números que sigan la siguiente lógica
Números en orden crecientes (del 1 en adelante)
Números que sean primos (1,2,3,5,7,9,11…)
Se pide implementar cada secuencia en una struct que siga cumpla con la interfaz

type IntSequence interface {
	Next() int
}

Por último se pide desarrollar un pequeño programa donde se pueda imprimir los primeros 30 números de una de las dos secuencias según una “tirada de dados”.
Tip: Podemos simular una tirada de dado usando rand.Intn(2) (del paquete math/rand).
*/
