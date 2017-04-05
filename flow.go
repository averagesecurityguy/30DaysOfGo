/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Demonstrate code flow in Go

Example usage:

$ go run flow.go
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	str := "This is a string."
	ba := []byte("This is another string.")

	// A string is a collection of immutable, UTF-8 encode, bytes. We can
	// iterate over the items in the string and print their index and their
	// value.

	fmt.Println("Print each index and char of a string with for loop.")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d: %c\n", i, str[i])
	}
	fmt.Println()

	// The more idiomatic way to do this would be to use the range keyword.
	fmt.Println("Print each index and char of a string with range loop.")
	for i, c := range str {
		fmt.Printf("%d: %c\n", i, c)
	}
	fmt.Println()

	// Either method could be used for byte arrays as well.
	fmt.Println("Use the range method to print each index and byte value.")
	for i, b := range ba {
		fmt.Printf("%d: %d\n", i, b)
	}
	fmt.Println()

	// If you only want the index or the value you can use an underscore to
	// discard the value you don't want.
	fmt.Println("Only get the characters in the string.")
	for _, c := range str {
		fmt.Println(c)
	}
	fmt.Println()

	fmt.Println("Only get the index values in the string.")
	for i, _ := range str {
		fmt.Println(i)
	}
	fmt.Println()

	// Loops can be broken with the break keyword and processing can continue
	// using the continue keyword.
	fmt.Println("Print only the odd characters up to 12.")
	for i, c := range str {
		if (i % 2) == 0 {
			continue
		} else if i >= 12 {
			break
		} else {
			fmt.Println(i, string(c))
		}
	}
	fmt.Println()

	// Go does not have a while loop. Use a for loop with appropriate index
	// values or use the range keyword. If you need an infinite loop you can
	// use the following syntax.
	fmt.Println("Infinite Loop")
	for {
		fmt.Println("Hit Ctrl-C to quit.")
		time.Sleep(2 * time.Second)
	}

}
