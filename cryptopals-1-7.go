/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Complete challenge 6 from Cryptopals set 1.
http://cryptopals.com/sets/1

Usage:

$ go run cryptopals-1-6.go
*/

package main

import (
    "os"
    "fmt"
    "bufio"
    "crypto/aes"
    "encoding/base64"
)


func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}


func get_data(filename string) []byte {
    file, err := os.Open(filename)
    check(err)

    var temp []byte
    scan := bufio.NewScanner(file)
	for scan.Scan() {
        temp = append(temp, []byte(scan.Text())...)
    }

    data := make([]byte, base64.StdEncoding.DecodedLen(len(temp)))
    _, err = base64.StdEncoding.Decode(data, temp)
    check(err)

    return data
}


func chunk(data []byte, size int) [][]byte {
    var chunks [][]byte

    for i:=0; i<len(data); i=i+size {
        chunks = append(chunks, data[i:i+size])
    }

    return chunks
}

func pad(data []byte, size int) []byte {
    if len(data) < size {
        pad := size - len(data)
        for i:=0; i<pad; i++ {
            data = append(data, byte(pad))
        }
    }

    return data
}


func main() {
    fmt.Println("Cryptopals Set 1")
    fmt.Println("================")

    fmt.Println("Challenge 7")
    fmt.Println("-----------")

    block_size := 16
    ecb, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
    check(err)

    ciphertext := get_data("data/7.txt")
    plaintext := make([]byte, 0)
    chunks := chunk(ciphertext, block_size)

    for _, chunk := range chunks {
        chunk = pad(chunk, block_size)
        temp := make([]byte, block_size)

        ecb.Decrypt(temp, chunk)

        plaintext = append(plaintext, temp...)
    }

    fmt.Println(string(plaintext))
}
