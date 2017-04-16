/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Gather any information Gravatar has about a particular email address.

Usage:

$ go run gravatar.go email_address
*/

package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
	}
}

func get(url string) []byte {
	resp, err := http.Get(url)
	check(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	if resp.StatusCode == 200 {
		return body
	} else {
		fmt.Printf("Document not available: %d\n", resp.StatusCode)
		return []byte("")
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run gravatar.go email_address")
		os.Exit(1)
	}

	email := string(os.Args[1])
	email = strings.ToLower(email)
	email = strings.Trim(email, " ")
	hash := fmt.Sprintf("%x", md5.Sum([]byte(email)))

	/*
	   Get Information From Gravatar.
	*/
	fmt.Printf("Getting information for: %s\n", email)
	url := fmt.Sprintf("https://www.gravatar.com/%s.json", hash)
	data := get(url)

	/*
	   Parse JSON Response
	*/
	var root map[string][]interface{}
	err := json.Unmarshal(data, &root)
	check(err)

	for _, entry := range root["entry"] {
		e := entry.(map[string]interface{})
		fmt.Printf("    Profile Url: %s\n", e["profileUrl"])
		fmt.Printf("    Preferred Username: %s\n", e["preferredUsername"])
		fmt.Printf("    Display Name: %s\n", e["displayName"])

		photos := e["photos"].([]interface{})
		for i, photo := range photos {
			p := photo.(map[string]interface{})
			fmt.Printf("    Photo %d: %s\n", i+1, p["value"])
		}

		switch name := e["name"].(type) {
		case map[string]interface{}:
			fmt.Printf("    Name: %s\n", name["formatted"])
		default:
			// Do nothing
		}

		fmt.Println()
	}
}
