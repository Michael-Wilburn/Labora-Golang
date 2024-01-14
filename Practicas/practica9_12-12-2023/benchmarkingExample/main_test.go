package main

import "testing"

func BenchmarkPrintWhenIsSaturdaySwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printWhenIsSaturdaySwitch()
	}
}
func BenchmarkPrintWhenIsSaturdayIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printWhenIsSaturdayIf()
	}
}
