/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Demonstrate strings and byte arrays in Go

Usage:

$ go run strings.go string
*/

// Let Go know that this file is the starting point for our script. It holds
// the main function.
package main

// Import any libraries needed.
import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run strings.go string")
		os.Exit(1)
	}

	// Go makes a distinction between strings and byte arrays. A string is
	// read only and is UTF-8 encoded. A byte array is an array of bytes.
	str := os.Args[1]

	// To check whether a variable is a string or a byte array at run time,
	// you can use the reflect package.
	fmt.Printf("str is type: %s\n", reflect.TypeOf(str))

	// A string can be converted to a byte (uint8) array.
	b := []byte(str)
	fmt.Printf("b is type: %s\n", reflect.TypeOf(b))

	// A byte array can be converted to a string
	s := string(b)
	fmt.Printf("s is type: %s\n", reflect.TypeOf(s))

	// The strings package allows you to manipulate strings in many ways. It
	// is way too much to cover in here so check out the details at
	// https://golang.org/pkg/strings/

	// You can write strings to the console using the fmt package. The Println
	// function will print the string you provide and add a newline character.
	fmt.Println("Printing a static string with fmt.Println.")
	s = "Printing a string variable with fmt.Println."
	fmt.Println(s)

	// You can also use the fmt.Printf function to print strings that contain
	// variables. Similar to the C printf and scanf functions.
	s = "My new string."
	fmt.Printf("This is my new string: %s\n", s)

	// A byte array can be interpreted as a string using Printf.
	b = []byte("Printing a []byte variable with fmt.Printf.")
	fmt.Printf("%s\n", b)

	// If you want to see the actual bytes (uint8 values) in the array use
	// fmt.Println.
	fmt.Println(b)

	// You can find all the details here, https://golang.org/pkg/fmt/, but to
	// get you started, %d is an integer, %s is a string, %x is hex.
	fmt.Printf("Integer: %d\n", 20)
	fmt.Printf("Hex: %x\n", 20)
	fmt.Printf("%x\n", b)

	// If you want to build a string instead of printing it out, you can use
	// the fmt.Sprintf function the same way you use the fmt.Printf function.
	s = fmt.Sprintf("Integer: %d", 20)
	fmt.Println(s)
	s = fmt.Sprintf("Hex: %x", 20)
	fmt.Println(s)
}
