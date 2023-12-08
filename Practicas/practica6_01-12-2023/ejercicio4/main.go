// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"math/rand"
)

type NumberGenerator interface {
	Next(n int) int
}

type AscendingNumberGenerator struct{}

func (_ AscendingNumberGenerator) Next(n int) int {
	return n + 1
}

type DescendingNumberGenerator struct{}

func (_ DescendingNumberGenerator) Next(n int) int {
	return n - 1
}

type NumberPredicate interface {
	Fulfill(n int) bool
}

type Primes struct{}

func (_ Primes) Fulfill(n int) bool {
	return isPrime(n)
}

func isPrime(n int) bool {
	if n <= 0 {
		n = -n // pequeña vuelta de tuerca para que funcione con números negativos. Tomo su valor como si fuera positivo!
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

type Integers struct{}

func (_ Integers) Fulfill(n int) bool {
	return true
}

type Evens struct{}

func (_ Evens) Fulfill(n int) bool {
	return n%2 == 0
}

type Sequence struct {
	current         int
	numberGenerator NumberGenerator
	numberPredicate NumberPredicate
}

func (s *Sequence) Next() int {
	var next int
	for n := s.numberGenerator.Next(s.current); n <= math.MaxInt32; n = s.numberGenerator.Next(n) {
		if s.numberPredicate.Fulfill(n) {
			next = n
			break
		}
	}
	s.current = next
	return next
}

func main() {
	var numberPredicate NumberPredicate
	random := rand.Intn(3)
	switch random {
	case 0:
		numberPredicate = Primes{}
	case 1:
		numberPredicate = Integers{}
	case 2:
		numberPredicate = Evens{}
	}

	var numberGenerator NumberGenerator
	random = rand.Intn(2)
	switch random {
	case 0:
		numberGenerator = AscendingNumberGenerator{}
	case 1:
		numberGenerator = DescendingNumberGenerator{}
	}

	var sequence Sequence = Sequence{current: 0, numberGenerator: numberGenerator, numberPredicate: numberPredicate}
	for i := 1; i <= 26; i++ {
		fmt.Println(sequence.Next())
	}

}
