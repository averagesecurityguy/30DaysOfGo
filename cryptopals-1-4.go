/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Complete challenge 4 from Cryptopals set 1.
http://cryptopals.com/sets/1

Usage:

$ go run cryptopals-1-4.go
*/

package main

import (
    "os"
    "fmt"
    "math"
    "bufio"
    "strings"
	"encoding/hex"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}


func from_hex(data string) []byte {
    b, err := hex.DecodeString(data)
    check(err)
    return b
}


func xor(b1 []byte, b2 byte) string {
    result := make([]byte, len(b1))

    for i, _ := range b1 {
        result[i] = b1[i] ^ b2
    }

    return string(result)
}


func score(data string) float64 {
    /*
    Score the string as English using chi-squared. We send a lower-cased string
    to this function so only have to worry about counting lowercase letters.
    */
    counts := make(map[rune]int)
    chi2 := 0.0
    alpha := "abcdefghijklmnopqrstuvwxyz "
    total := 0
    freq := map[rune]float64 {
        'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253, 'e': 0.12702,
        'f': 0.02228, 'g': 0.02015, 'h': 0.06094, 'i': 0.06966, 'j': 0.00153,
        'k': 0.00772, 'l': 0.04025, 'm': 0.02406, 'n': 0.06749, 'o': 0.07507,
        'p': 0.01929, 'q': 0.00095, 'r': 0.05987, 's': 0.06327, 't': 0.09056,
        'u': 0.02758, 'v': 0.00978, 'w': 0.02360, 'x': 0.00150, 'y': 0.01974,
        'z': 0.00074, ' ': 0.23200}

    // Get a count of all the letters and the total number of letters.
    for _, c := range data {
        if strings.Contains(alpha, string(c)) {
            total = total + 1
            _, ok := counts[c]
            if ok {
                counts[c] = counts[c] + 1
            } else {
                counts[c] = 1
            }
        }
    }

    // Do not calculate the chi-squared value unless the string is at least 70%
    // ASCII alphabet.
    if total < int(float64(0.7) * float64(len(data))) {
        return 1000.0
    }

    // Calculate chi-squared for each letter
    for _, k := range alpha {
        expected := float64(total) * freq[k]
        actual := float64(counts[k])
        val := math.Pow(actual - expected, 2)/expected
        chi2 = chi2 + val
    }

    return chi2
}


func break_xor(e string) (float64, string) {
    low := 1000.0
    msg := ""

    // Bruteforce the key by XORing each possible key, analyzing the decrypted
    // message, and scoring it. Lowest score wins.
    for i:=0; i<256; i++ {
        k := byte(i)
        dec := xor(from_hex(e), k)
        total := score(strings.ToLower(dec))

        if total < low {
            low = total
            msg = string(dec)
        }
    }

    return low, msg
}


func open(filename string) *os.File {
	/*
	   Open a file as read only.
	*/
	data, err := os.Open(filename)
	check(err)

	return data
}


func main() {
    fmt.Println("Cryptopals Set 1")
    fmt.Println("================")

    fmt.Println("Challenge 4")
    fmt.Println("-----------")
    crypts := open("data/4.txt")
    cscan := bufio.NewScanner(crypts)

	for cscan.Scan() {
        enc := cscan.Text()
		score, msg := break_xor(enc)
        if score < 35.0 {
            fmt.Printf("%q\n", msg)
        }
    }
}
