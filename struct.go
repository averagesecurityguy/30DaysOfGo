/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Demonstrate Structs in Go

Usage:

$ go run struct.go
*/

package main

import (
	"fmt"
)

// Go has a number of built in types, which are useful for most situations.
// Sometimes, though, you need a more complex type. Go provides Structs for
// these situations. Structs are declared using the following syntax.
type Address struct {
	street1 string
	street2 string
	city    string
	state   string
	zip     string
}

// The members of the Struct can be either a built in type or another Struct.
type User struct {
	id      int
	first   string
	last    string
	address Address
}

func print_user(u User) {
	fmt.Printf("%s %s (%d)\n", u.first, u.last, u.id)
	fmt.Println(u.address.street1)
	if u.address.street2 != "" {
		fmt.Println(u.address.street2)
	}
	fmt.Printf("%s, %s  %s\n", u.address.city, u.address.state, u.address.zip)
	fmt.Println("")
}

func main() {

	// Create a new User and populate it.
	var user1 User

	user1.id = 1
	user1.first = "Average"
	user1.last = "SecurityGuy"
	user1.address.street1 = "123 Some Street"
	user1.address.street2 = ""
	user1.address.city = "Some City"
	user1.address.state = "AZ"
	user1.address.zip = "55555"

	// Println will print our Struct but it doesn't look very nice.
	fmt.Println("Print our User struct.")
	fmt.Println(user1)
	fmt.Println("")

	// We can use our own function to print the Struct
	fmt.Println("Pretty print our user.")
	print_user(user1)

	// We can modify our User as well.
	fmt.Println("Print updated user.")
	user1.address.street2 = "Apt. 100"
	print_user(user1)

	// We can also make an array of Users and act upon each User in the array.
	var users []User
	users = append(users, user1)

	fmt.Println("Printing a list of users.")
	for _, u := range users {
		print_user(u)
	}
}
