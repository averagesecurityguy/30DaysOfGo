/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Encrypt or decrypt a file using a passphrase. The .enc file extension is added
to the encrypted file and removed after decryption. The original file is
overwritten during encryption or decryption.

Usage:

Encrypt file
$ go run krypt.go -f filename -p passphrase

Decrypt file
$ go run krypt.go -d -f filename.enc -p passphrase

*/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"flag"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		stop(e.Error())
	}
}

func stop(msg string) {
	fmt.Printf("Error: %s\n", msg)
	os.Exit(1)
}

func read_file(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	check(err)

	return data
}

func write_file(data []byte, filename string) {
	err := ioutil.WriteFile(filename, data, 0640)
	check(err)
}

func save(data []byte, new_fn, old_fn string) {
	write_file(data, new_fn)
	err := os.Remove(old_fn)
	check(err)
}

func random_bytes(count int) []byte {
	bytes := make([]byte, count)
	_, err := io.ReadFull(rand.Reader, bytes)
	check(err)

	return bytes
}

func encrypt(data, key []byte) []byte {
	// Modified from https://golang.org/pkg/crypto/cipher/#example_NewGCM_encrypt
	block, err := aes.NewCipher(key)
	check(err)

	nonce := random_bytes(12)

	aesgcm, err := cipher.NewGCM(block)
	check(err)

	ciphertext := aesgcm.Seal(nil, nonce, data, nil)

	return append(nonce, ciphertext...)
}

func decrypt(data, key []byte) []byte {
	// Modified from https://golang.org/pkg/crypto/cipher/#example_NewGCM_decrypt
	nonce := data[:12]
	ciphertext := data[12:]

	block, err := aes.NewCipher(key)
	check(err)

	aesgcm, err := cipher.NewGCM(block)
	check(err)

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	check(err)

	return plaintext
}

func main() {
	/*
	   Parse and Verify Arguments
	*/
	// Link to flag module documentation. Brief explanation of code.
	var d bool
	var filename string
	var phrase string

	flag.BoolVar(&d, "d", false, "Decrypt a file.")
	flag.StringVar(&filename, "f", "", "File to be encrypted or decrypted.")
	flag.StringVar(&phrase, "p", "", "Passphrase for encryption/decryption.")
	flag.Parse()

	if len(phrase) < 20 {
		flag.Usage()
		stop("The passphrase must be at least 20 characters.")
	}

	/*
	   Generate Key From Passphrase
	*/
	// The first 12 characters of the passphrase will be used as the salt.
	phrase_byte := []byte(phrase)
	salt := []byte(phrase[:12])
	key := pbkdf2.Key(phrase_byte, salt, 4096, 32, sha256.New)

	/*
	   Encrypt or Decrypt File
	*/
	data := read_file(filename)

	if d {
		new_filename := filename[:len(filename)-4]
		new_data := decrypt(data, key)
		save(new_data, new_filename, filename)
	} else {
		new_filename := fmt.Sprintf("%s.enc", filename)
		new_data := encrypt(data, key)
		save(new_data, new_filename, filename)
	}
}
