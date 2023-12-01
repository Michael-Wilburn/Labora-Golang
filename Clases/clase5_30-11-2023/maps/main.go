package main

import "fmt"

func main() {
	var isPrimeByInt map[int]bool = make(map[int]bool)

	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			isPrimeByInt[i] = true
		}
	}

	fmt.Println(isPrimeByInt)
	isPrimeByInt2 := isPrimeByInt
	isPrimeByInt2[1] = false
	fmt.Println(isPrimeByInt2)

	planetCountBySystem := map[string]int{
		"Sol":      9,
		"Centauri": 2,
		"Loner":    0,
	}
	planetCount, exists := planetCountBySystem["Loner"]
	fmt.Println(planetCount, exists)

}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
