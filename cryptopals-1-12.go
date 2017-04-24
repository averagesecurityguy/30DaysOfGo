/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Complete the first two challenges from Cryptopals set 1.
http://cryptopals.com/sets/1

Usage:

$ go run cryptopals-1-12.go
*/

package main

import (
    "os"
    "fmt"
	"encoding/base64"
	"encoding/hex"
)


func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}


func from_b64(data string) []byte {
    bytes, err := base64.StdEncoding.DecodeString(data)
    check(err)
    return bytes
}


func to_b64(data []byte) string {
    return base64.StdEncoding.EncodeToString(data)
}


func from_hex(data string) []byte {
    bytes, err := hex.DecodeString(data)
    check(err)
    return bytes
}


func to_hex(data []byte) string {
    return hex.EncodeToString(data)
}


func xor_arrays(b1, b2 []byte) []byte {
    if len(b1) != len(b2) {
        panic("Cannot join byte arrays of two differentl lengths.")
    }

    result := make([]byte, len(b1))

    for i, _ := range b1 {
        result[i] = b1[i] ^ b2[i]
    }

    return result
}




func main() {
    fmt.Println("Cryptopals Set 1")
    fmt.Println("================")

    fmt.Println("Challenge 1")
    fmt.Println("-----------")
    hex_str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    bytes := from_hex(hex_str)
    b64 := to_b64(bytes)
    fmt.Println(b64)

    fmt.Println()
    fmt.Println("Challenge 2")
    fmt.Println("-----------")
    b1 := from_hex("1c0111001f010100061a024b53535009181c")
    b2 := from_hex("686974207468652062756c6c277320657965")
    fmt.Println(to_hex(xor_arrays(b1, b2)))

}
