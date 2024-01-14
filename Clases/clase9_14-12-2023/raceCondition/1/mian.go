package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter = 0
	const n = 10000
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Contador final:", counter)
}
