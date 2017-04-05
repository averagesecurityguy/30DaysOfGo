/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Demonstrate function usage in Go

Usage:

$ go run functions.go string
*/

package main

import (
	"fmt"
	"os"
)

// Declare a function that takes two parameters and does not return a value.
// The variable is declared before the type. Variables of the same type can be
// declared at the same time before stating the type:
// func some_func(a, b, c string, i, j int)
func write_string(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(s)
	}
}

// Declare a function that takes a string and returns a boolean value.
func short_string(s string) bool {
	return len(s) < 5
}

// Declare a function that returns two values. You can return as many value as
// needed by placing them in the comma delimited return list.
func test_string(s string) (int, bool) {
	l := len(s)
	return l, l < 5
}


func main() {

	// Check the number of arguments we have. The name of the script is the
	// first argument.
	if len(os.Args) != 2 {

		// If we don't have the correct number of arguments then print a
		// message. Println will automatically add a newline at the end of the
		// string. Always use double quotes for strings, single quotes have a
		// different meaning, which we are not going to discuss.
		fmt.Println("Usage: go run functions.go string")

		// Exit the program
		os.Exit(1)
	}

	// Declare and assign a variable. The alternative is to do it in two steps:
	// var str string
	// str = os.Args[1]
	str := os.Args[1]

	// Call a function with two parameters.
	write_string(str, 10)

	// Call a function and catch the return value.
	ans := short_string(str)
	if ans == true {
		fmt.Println("String is short")
	} else {
		fmt.Println("String is not short")
	}

	// Call a function and get both return values. If you don't need all of the
	// return values from a function, you can throw away the return values you
	// do not want or need. using an underscore:
	// _, b := test_string(str)
	n, b := test_string(str)
	if b == true {
		fmt.Printf("The string is short: %d characters.\n", n)
	} else {
		fmt.Printf("The string is not short: %d characters.\n", n)
	}
}
