/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Demonstrate goroutines and channels

Usage:

$ go run routine_channels.go
*/

package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
Go routines allow us to execute code asyncronously. Keep in mind that when the
main function ends then all of the child routines will be stopped. To prevent
this we have to wait for all of the child routines to finish using a WaitGroup.
*/

// Declare our WaitGroup globally so all functions can access it.
var wait sync.WaitGroup

func capitalize(word string) {
	// wait.Done() informs the WaitGroup that the go routine is complete. We
	// defer the call to wait.Done() until after the function completes.
	defer wait.Done()
	fmt.Println(strings.Title(word))
}

func upper(ch chan string) {
	// Loop until the channel is empty.
	for {
		// Pull a word from the channel
		word := <-ch

		// If word is "" then the channel is empty.
		if word == "" {
			break
		}

		fmt.Println(strings.ToUpper(word))
	}
}

func add(ch chan string, words []string) {
	for _, word := range words {
		ch <- word
	}
}

func main() {
	words := []string{"test1", "test2", "test3"}

	// wait.Add() tells our WaitGroup there is a routine it needs to wait on.
	// The parameter passed to Add is the number of routines that need to be
	// waited on.
	fmt.Println("Use go routines to capitalize our words.")
	wait.Add(len(words))

	for i := 0; i < len(words); i++ {
		go capitalize(words[i])
	}

	// wait.Wait() monitors the WaitGroup to ensure all go routines are
	// complete. Ensure a wait.Done() call is made for each wait.Add() call.
	wait.Wait()
	fmt.Println()

	// This code tells the WaitGroup that there are four routines to monitor
	// but the function tells the WaitGroup only three of the routines are
	// done. This will cause the script to crash. Uncomment the code if you
	// want to see this in action.
	/*
	   fmt.Println("Add more routines to the WaitGroup than we close.")
	   fmt.Println("This will crash.")
	   wait.Add(len(words) + 1)
	   for i:=0; i<len(words); i++ {
	       go capitalize(words[i])
	   }
	   wait.Wait()
	*/

	// Channels go hand in hand with go routines. For now, you can think of
	// channels as queues. Go routines can be used to put data on the queue
	// and to take data off the queue. You can have one or many go routines
	// feeding the channel and you can have one or more go routines being fed
	// by the channel. You can even have multiple channels being operated on
	// by multiple go routines. I will do my best to give you good examples
	// but you should also consult https://tour.golang.org/concurrency/2 and
	// https://gobyexample.com/channels and
	// http://guzalexander.com/2013/12/06/golang-channels-tutorial.html.

	fmt.Println("Use channels to upper case our words.")

	// Make a channel to operate on. This channel takes strings but you can
	// make channels that take ints, structs, etc.
	word_chan := make(chan string)

	// By default when you write to a channel the go script blocks until a read
	// is performed. We create a go routine that will read from the channel
	// and kick it off first. Then we write to the channel. If we tried to
	// write first, we would end up in a deadlock where each go routine is
	// waiting on the other. We do not have to manage waiting in this case. I'm
	// still trying to understand why.

	// Create a go routine to process the channel data.
	go upper(word_chan)

	// Add words to our channel.
	add(word_chan, words)

	close(word_chan)
}
