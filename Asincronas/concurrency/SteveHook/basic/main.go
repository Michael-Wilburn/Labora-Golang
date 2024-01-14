package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		fmt.Println("1")
		defer wg.Done()
	}()
	go func() {
		fmt.Println("2")
		defer wg.Done()
	}()
	go func() {
		fmt.Println("3")
		defer wg.Done()
	}()

	wg.Wait()
}
