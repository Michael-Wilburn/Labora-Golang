package main

import "fmt"

type MySlice []int

func (s MySlice) Remove(index int) MySlice {
	return append(s[:index], s[index+1:]...)
}

func main() {
	numbers := MySlice{1, 2, 3, 4, 5}
	// fmt.Println(removeFromSlice(numbers, 1))
	// fmt.Println(removeFromSliceWithOrder(numbers, 1))
	fmt.Println(numbers.Remove(1))

}

func removeFromSlice(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func removeFromSliceWithOrder(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
