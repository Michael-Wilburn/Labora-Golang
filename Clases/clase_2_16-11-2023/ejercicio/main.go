package main

import "fmt"

func main() {
	// var n int = 0
	// for n != 1 || n != 0 {
	// 	n++
	// }
	fmt.Println(calculateE(20))
}

func calculateE(loops float64) float64 {
	var ac, d, i float64
	ac = 0
	d = 1
	for i = 1; i <= loops; i++ {
		d = d * i
		ac = ac + 1/d
	}
	return ac + 1
}
