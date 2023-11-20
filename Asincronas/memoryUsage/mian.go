package main

import (
	"fmt"
	"runtime"
)

func main() {
	bitsPerNumber8()

	bitsPerNumber64()

}

func bitsPerNumber64() {
	var list []int
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)
}

func bitsPerNumber8() {
	var list []int8
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)
}
