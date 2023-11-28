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

	fmt.Printf("%+v\n", person)
	fmt.Printf("%+v", person.getIdPlusOne())
	// fmt.Printf("%+v", Person.getIdPlusOne(person))

}

func (person Person) getIdPlusOne() int {
	return person.Id + 1
}
