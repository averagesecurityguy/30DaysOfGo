/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Test which HTTP methods a Web Server or endpoint will support.

Usage:

$ go run http_method_check.go url_file
*/

package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func print_error(msg string) {
	fmt.Printf("[E] %s\n", msg)
}

func catch_panic() {
	if r := recover(); r != nil {
		fmt.Printf("[P] %s\n", r)
	}
}

func checkMethod(url, method string) {
	/*
	   Check the given method against the given URL. If we get a 200 then the
	   method is assumed implemented; anything else and the method is not.
	*/
	defer catch_panic()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		print_error(err.Error())
	}

	resp, err := client.Do(req)
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		print_error(err.Error())
	}

	if resp.StatusCode != 200 {
		fmt.Printf("The %s method on %s IS NOT implemented: %d.\n", method, url, resp.StatusCode)
	} else {
		fmt.Printf("The %s method on %s is implemented.\n", method, url)
	}
}

func open(filename string) *os.File {
	/*
	   Open a file as read only.
	*/
	data, err := os.Open(filename)
	if err != nil {
		print_error(err.Error())
	}

	return data
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run http_method_check.go url_file")
		os.Exit(1)
	}

	methods := []string{"ACL", "BASELINE-CONTROL", "BCOPY", "BDELETE",
		"BMOVE", "BPROPFIND", "BPROPPATCH", "CHECKIN",
		"CHECKOUT", "COPY", "DEBUG", "DELETE",
		"GET", "HEAD", "INDEX", "LABEL", "LOCK", "MERGE",
		"MKACTIVITY", "MKCOL", "MKDIR", "MKWORKSPACE", "MOVE",
		"NOTIFY", "OPTIONS", "ORDERPATCH", "PATCH", "POLL",
		"POST", "PROPFIND", "PROPPATCH", "PUT", "REPORT",
		"RMDIR", "RPC_IN_DATA", "RPC_OUT_DATA", "SEARCH",
		"SUBSCRIBE", "TRACE", "UNCHECKOUT", "UNLOCK", "UPDATE",
		"UNSUBSCRIBE", "VERSION-CONTROL", "X-MS-ENUMATTS"}

	urls := open(os.Args[1])
	fscan := bufio.NewScanner(urls)

	for fscan.Scan() {
		url := fscan.Text()

		if strings.HasPrefix(url, "#") == true {
			continue
		}

		for _, m := range methods {
			checkMethod(url, m)
		}
	}

	defer urls.Close()
}
