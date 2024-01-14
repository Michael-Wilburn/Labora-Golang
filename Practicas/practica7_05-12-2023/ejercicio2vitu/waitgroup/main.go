// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"time"
)

type SumResult struct {
	Type  string
	Value int
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	var oddsSumDuration, evensSumDuration time.Duration
	go func() {
		begins := time.Now()
		sum := sumFirst100Evens()
		ends := time.Now()
		evensSumDuration = ends.Sub(begins)
		fmt.Println("Suma de los primeros 100 pares es", sum)
		wg.Done()
	}()
	go func() {
		begins := time.Now()
		sum := sumFirst100Odds()
		ends := time.Now()
		oddsSumDuration = ends.Sub(begins)
		fmt.Println("Suma de los primeros 100 impares", sum)
		wg.Done()
	}()

	fmt.Println("Esperando....")
	wg.Wait()
	fmt.Println("evensSumLapse:", evensSumDuration)
	fmt.Println("oddsSumLapse:", oddsSumDuration)
	if evensSumDuration < oddsSumDuration {
		fmt.Println("Gano par")
	} else {
		fmt.Println("Gano impar")
	}
	fmt.Println("Fin feliz")
}

func sumFirst100Evens() int {
	sum := 0
	for i := 1; i <= 10000000; i++ {
		sum += (i * 2)
	}
	return sum
}

func sumFirst100Odds() int {
	sum := 0
	for i := 1; i <= 10000000; i++ {
		sum += (i * 2) - 1
	}
	return sum
}
