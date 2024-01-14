// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

type SumResult struct {
	Type     string // "odd" o "even"
	Sum      int
	Duration time.Duration
}

func main() {
	var sumResultChannel chan SumResult = make(chan SumResult)

	go func() {
		begins := time.Now()
		sum := sumFirst100Evens()
		ends := time.Now()
		duration := ends.Sub(begins)
		sumResult := SumResult{
			Type:     "Evens",
			Sum:      sum,
			Duration: duration,
		}
		sumResultChannel <- sumResult
	}()
	go func() {
		begins := time.Now()
		sum := sumFirst100Odds()
		ends := time.Now()
		duration := ends.Sub(begins)
		sumResult := SumResult{
			Type:     "Odds",
			Sum:      sum,
			Duration: duration,
		}
		sumResultChannel <- sumResult
	}()

	sumResultOne := <-sumResultChannel
	sumResultTwo := <-sumResultChannel

	fmt.Println("Esperando....")
	fmt.Printf("%+v\n%+v\n", sumResultOne, sumResultTwo)
	if sumResultOne.Duration < sumResultTwo.Duration {
		fmt.Println("Gano", sumResultOne.Type)
	} else {
		fmt.Println("Gano", sumResultTwo.Type)
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
