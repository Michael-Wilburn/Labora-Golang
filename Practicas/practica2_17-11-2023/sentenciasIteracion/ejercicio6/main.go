package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fizzBuzz()
	fizzBuzzAlternative()
}

//6.Ejercicio FizzBuzz en su forma clásica es el siguiente: "Para cada número del 1 al 100, imprime 'Fizz' si el número es divisible por 3, 'Buzz' si es divisible por 5 y 'FizzBuzz' si es divisible por ambos. Si no es divisible ni por 3 ni por 5, simplemente imprime el número."

func fizzBuzz() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// Otra forma de hacerlo
func fizzBuzzAlternative() {
	for i := 1; i <= 100; i++ {
		out := ""
		if i%3 == 0 {
			out += "Fizz"
		}
		if i%5 == 0 {
			out += "Buzz"
		}
		if out == "" {
			out = strconv.Itoa(i)
		}
		fmt.Println(out)
	}
}
