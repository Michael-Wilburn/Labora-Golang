package main

import "fmt"

var users = []string{}

func addUser(user ...string) {
	users = append(users, user...)
}

func main() {
	addUser("Alice", "Bob", "Foo", "Bar")

	fmt.Println(users)
}
