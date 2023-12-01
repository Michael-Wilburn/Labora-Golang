package main

import "fmt"

/*
When
1. When we need to update a state
pointer = 8 bytes

*/

type User struct {
	email    string
	username string
	age      int
}

func (u User) Email() string {
	return u.email
}

func (u *User) UpdateEmail(email string) {
	u.email = email
}

func Email(user User) string {
	return user.email
}

func main() {
	user := User{
		email: "agg@foo.com",
	}
	user.UpdateEmail("mike@cool.com")
	fmt.Println(user.Email())
}
