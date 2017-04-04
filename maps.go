/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Demonstrate maps (keyed lists or dictionaries) in Go

Usage:

$ go run maps.go
*/

package main

import (
    "fmt"
    "crypto/md5"
)

func main() {
    // A map is a key value store or a dictionary (if you come from Python).
    // The key can be any type that is comparable and the value can be any
    // type. You can get more details about maps here:
    // https://blog.golang.org/go-maps-in-action.
    //
    // We will use only built in types as keys because we know they are
    //comparable. Declare a map using the following syntax:
    // m := make(map[keyType]valType)
    m := make(map[string]string)

    // We now have a variable that maps one string to another string.
    m["test"] = fmt.Sprintf("%x", md5.Sum([]byte("test")))
    fmt.Println("Print our map.")
    fmt.Println(m)
    fmt.Println("")

    // Add more values to our map using a for loop.
    for i:=0; i<10; i++ {
        s := fmt.Sprintf("test%d", i)
        m[s] = fmt.Sprintf("%x", md5.Sum([]byte(s)))
    }

    fmt.Println("Print our expanded map.")
    fmt.Println(m)
    fmt.Println("")

    fmt.Println("Pretty print our map.")
    // Use range to pretty print the map
    for k, v := range m {
        fmt.Printf("%s: %s\n", k, v)
    }
    fmt.Println("")

    // Map values can be updated
    m["test"] = ""

    // Map keys can be deleted
    delete(m, "test0")

    fmt.Println("Print our updated map.")
    for k, v := range m {
        fmt.Printf("%s: %s\n", k, v)
    }
    fmt.Println("")

    // Assign a value from the map to a variable
    h := m["test1"]

    // This can cause a problem if the key does not exist so let's check to see
    // if the key exists while making the assignment.
    h, ok := m["test10"]
    if ok {
        fmt.Println(h)
    } else {
        fmt.Println("This key does not exist.")
    }
}
