/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Encode a string using base64, base32, hex (ff), or C hex (\xff).

Usage:

$ go run encode.go base64 Test
VGVzdA==
$ go run encode.go base32 Test
KRSXG5A=
$ go run encode.go hex Test
54657374
$ go run encode.go c Test
\x54\x65\x73\x74
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
	fmt.Println("Usage: go run encode.go format string")
	fmt.Println("Accepted formats are: base64, base32, hex, or c")
	os.Exit(0)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	format := strings.ToLower(os.Args[1])
	data := []byte(os.Args[2])

	switch format {
	case "base64":
		fmt.Println(base64.StdEncoding.EncodeToString(data))
	case "base32":
		fmt.Println(base32.StdEncoding.EncodeToString(data))
	case "hex":
		fmt.Println(hex.EncodeToString(data))
	case "c":
		for _, b := range data {
			fmt.Printf("\\x%x", b)
		}
		fmt.Println("")
	default:
		usage()
	}
}
