package main

import "testing"

func BenchmarkSumaSecuencial(b *testing.B) {
	var array []int
	for i := 0; i <= 10000000; i++ {
		array = append(array, i)
	}
	for i := 0; i < b.N; i++ {
		sumaSecuencial(array)
	}
}
func BenchmarkSumaConcurrente(b *testing.B) {
	var array []int
	for i := 0; i <= 10000000; i++ {
		array = append(array, i)
	}
	for i := 0; i < b.N; i++ {
		sumaConcurrente(array)
	}
}
