/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Finds all pastes listed on Have I Been Pwnd for the specified email and
downloads each identified paste if it is available.

Usage:

$ go run hibp_paste.go email_address
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type Paste struct {
	Source string
	Id     string
}

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
	}
}

func get(url string) []byte {
	fmt.Println(url)

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
		fmt.Println("Usage: go run hibp_paste.go email_address")
		os.Exit(1)
	}

	email := os.Args[1]
	re := regexp.MustCompile(fmt.Sprintf(".*%s.*", email))

	/*
	   Get Paste List from HIBP.
	*/
	url := fmt.Sprintf("https://haveibeenpwned.com/api/v2/pasteaccount/%s", email)
	data := get(url)

	var pastes []Paste
	json.Unmarshal(data, &pastes)

	/*
	   Download Each Paste
	*/
	for _, p := range pastes {
		var url string
		var data []byte

		switch p.Source {
		case "Pastebin":
			url = fmt.Sprintf("https://pastebin.com/raw/%s", p.Id)
			data = get(url)
		case "Slexy":
			url = fmt.Sprintf("http://slexy.org/raw/%s", p.Id)
			data = get(url)
		default:
			fmt.Printf("Paste source does not support raw viewing - %s: %s\n", p.Source, p.Id)
			data = []byte("")
		}

		// Print only the data that matches our email.
		matches := re.FindAllString(string(data), -1)

		if len(matches) > 0 {
			fmt.Println(strings.Join(matches, "\n"))
		}

		fmt.Println("")
		time.Sleep(10 * time.Second)
	}
}
