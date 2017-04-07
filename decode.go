/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Decode a string using base64, base32, hex (ff), or C hex (\xff).

Usage:

$ go run decode.go base64 VGVzdA==
Test
$ go run decode.go base32 KRSXG5A=
Test
$ go run decode.go hex 54657374
Test
$ go run decode.go c "\x54\x65\x73\x74"
Test
*/

package main

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: go run decode.go format string")
	fmt.Println("Accepted formats are: base64, base32, hex, or c")
	os.Exit(0)
}

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	format := strings.ToLower(os.Args[1])
	data := os.Args[2]

	switch format {
	case "base64":
		bytes, err := base64.StdEncoding.DecodeString(data)
		check(err)
		fmt.Println(string(bytes))
	case "base32":
		bytes, err := base32.StdEncoding.DecodeString(data)
		check(err)
		fmt.Println(string(bytes))
	case "hex":
		bytes, err := hex.DecodeString(data)
		check(err)
		fmt.Println(string(bytes))
	case "c":
		bytes := strings.Split(data, "\\x")
		for _, b := range bytes {
			b, err := hex.DecodeString(b)
			check(err)
			fmt.Printf("%s", b)
		}
		fmt.Println("")
	default:
		usage()
	}
}
