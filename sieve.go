/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Use the Sieve of Eratosthanes to find the number of primes up to a number n.

Example usage:

$ go run sieve.go number
*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// A simple implementation of the Sieve of Eratosthanes
// Modified from: https://gist.githubusercontent.com/dazfuller/3940659/raw/f9ef272aea070860d9cfe67715b5b401de132745/solution3_snippet.go
func Sieve(n uint64) []uint64 {
	// If n is less than 2 return an empty array
	if n < uint64(2) {
		return make([]uint64, 0)
	}

	// Go defaults a bool array to false
	sieve := make([]bool, n)
	sieve[0] = true
	sieve[1] = true

	limit := uint64(math.Sqrt(float64(n))) + uint64(1)

	// Generate the sieve by removing multiples of primes
	lp := uint64(2)
	for lp < limit {
		for i := lp * 2; i < n; i += lp {
			sieve[i] = true
		}

		// Find next prime candidate. Will be between last prime and lp*2
		for i := lp + 1; i < lp*2; i++ {
			if sieve[i] == false {
				lp = i
				break
			}
		}
	}

	// Count the number of primes in the sieve
	count := 0
	for _, v := range sieve {
		if v == false {
			count++
		}
	}

	// Build the prime list by looking for the primes in the sieve
	result := make([]uint64, count)
	index := uint64(0)
	for i, v := range sieve {
		if v == false {
			result[index] = uint64(i)
			index++
		}
	}

	return result
}

func usage() {
	fmt.Println("Usage: go run sieve.go number")
	os.Exit(0)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		usage()
	}

	primes := Sieve(n)

	fmt.Printf("There are %d primes less than %d.\n", len(primes), n)
}
