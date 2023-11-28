package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func main() {
	var person Person

	person.Id = 2
	person.Name = "pepe"

	fmt.Println(person.Id, person.Name)

}
