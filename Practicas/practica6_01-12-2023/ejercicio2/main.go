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
	Title() string
}

type LinealSequence struct {
	current int
}

func (s *LinealSequence) Next() int {
	var next int = s.current
	s.current++
	return next
}

func (_ LinealSequence) Title() string {
	return "Secuencia lineal (creciente)"
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

func (_ PrimesSequence) Title() string {
	return "Secuencia de primos (creciente)"
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
	fmt.Println(intSequence.Title())
	for i := 1; i <= 50; i++ {
		fmt.Println(intSequence.Next())
	}

}
