package main

import (
	"fmt"
	"sort"
)

type Numbers []int

func (n Numbers) Len() int { return len(n) }
func (n Numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n Numbers) Less(i, j int) bool {
	return n[i] < n[j]
}

type byDec struct {
	Numbers
}

func (n byDec) Len() int { return len(n.Numbers) }
func (n byDec) Swap(i, j int) {
	n.Numbers[i], n.Numbers[j] = n.Numbers[j], n.Numbers[i]
}
func (n byDec) Less(i, j int) bool {
	return n.Numbers[i] > n.Numbers[j]
}

type byInc struct {
	Numbers
}

func (n byInc) Len() int { return len(n.Numbers) }
func (n byInc) Swap(i, j int) {
	n.Numbers[i], n.Numbers[j] = n.Numbers[j], n.Numbers[i]
}
func (n byInc) Less(i, j int) bool {
	return n.Numbers[i] < n.Numbers[j]
}

func main() {
	numbers := Numbers{1, 10, 4, 9, 3}
	sort.Sort(byInc{numbers})
	fmt.Println(numbers)

	sort.Sort(byDec{numbers})
	fmt.Println(numbers)
}
