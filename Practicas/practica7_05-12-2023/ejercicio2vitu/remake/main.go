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

var sumResultChannel chan SumResult = make(chan SumResult)

func main() {
	go sumFirst100(sumFirst100Evens, "evens")
	go sumFirst100(sumFirst100Odds, "odds")

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

type FuncSumFirst100 func() int

func sumFirst100(sumFunc FuncSumFirst100, _type string) {
	begins := time.Now()
	sum := sumFunc()
	ends := time.Now()
	duration := ends.Sub(begins)
	sumResult := SumResult{
		Type:     _type,
		Sum:      sum,
		Duration: duration,
	}
	sumResultChannel <- sumResult
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
