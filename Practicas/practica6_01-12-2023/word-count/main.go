package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I ate a donut. Then I ate another donut."
	fmt.Println(WordCount(s))
}

func WordCount(s string) map[string]int {
	var ocurrencesByChar map[string]int = make(map[string]int)
	v := strings.Fields(s)
	for _, char := range v {
		_, exists := ocurrencesByChar[string(char)]
		if !exists {
			ocurrencesByChar[string(char)] = 0
		}
		ocurrencesByChar[string(char)]++
	}

	return ocurrencesByChar
}
