/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Hash a string using MD5, SHA1, SHA256, or SHA512.

Usage:

$ go run hash.go md5 Test
0cbc6611f5540bd0809a388dc95a615b
$ go run hash.go sha1 Test
640ab2bae07bedc4c163f679a746f7ab7fb5d1fa
$ go run hash.go sha256 Test
532eaabd9574880dbf76b9b8cc00832c20a6ec113d682299550d7a6e0f345e25
$ go run hasg.go sha512 Test
c6ee9e33cf5c6715a1d148fd73f7318884b41adcb916021e2bc0e800a5c5dd97f5142178f6ae88c8fdd98e1afb0ce4c8d2c54b5f37b30b7da1997bb33b0b8a31
*/

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: go run hash.go hash string")
	fmt.Println("Accepted hashes are: md5, sha1, sha256, or sha512")
	os.Exit(0)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	hash := strings.ToLower(os.Args[1])
	data := []byte(os.Args[2])

	switch hash {
	case "md5":
		fmt.Printf("%x\n", md5.Sum(data))
	case "sha1":
		fmt.Printf("%x\n", sha1.Sum(data))
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256(data))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512(data))
	default:
		usage()
	}
}
