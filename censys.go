/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

View Censys data about an IP address, a domain name or a TLS certificate.

Usage:

$ go run censys.go -d domain
$ go run censys.go -i ip
$ go run censys.go -c sha256
*/
package main

import (
	"os"
    "fmt"
    "flag"
    "bytes"
    "net/http"
    "io/ioutil"
	"encoding/json"
)


const base = "https://censys.io/api/v1/view"
const uid = ""
const secret = ""


func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
	}
}


func lookup(data_type, value string) {
    url := fmt.Sprintf("%s/%s/%s", base, data_type, value)
    fmt.Printf("[*] Looking up %s\n", url)

    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    req.SetBasicAuth(uid, secret)
    resp, err := client.Do(req)
	check(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	if resp.StatusCode == 404 {
        fmt.Println("[-] File not found.")
    } else if resp.StatusCode == 429 {
        fmt.Println("[-] Rate limit exceeded. Please wait and try again.")
    } else if resp.StatusCode == 500 {
        fmt.Println("[-] Internal Server Error.")
    } else {
        var pp bytes.Buffer
        err := json.Indent(&pp, body, "", "  ")
        check(err)

        fmt.Println(string(pp.Bytes()))
	}
}


func main() {
	var d string
	var i string
	var c string

	flag.StringVar(&d, "d", "", "Domain Name")
	flag.StringVar(&i, "i", "", "IP Address")
	flag.StringVar(&c, "c", "", "SHA256 Fingerprint")

	flag.Parse()

    // We should only have one flag set
    if flag.NFlag() != 1 {
        fmt.Println("[-] Only one of the flags -d, -i, or -c may be set.")
        fmt.Println("[-] Use -h or --help for more details.")
        os.Exit(0)
    }

	if d != "" {
        lookup("websites", d)
    } else if i != "" {
        lookup("ipv4", i)
    } else {
        lookup("certificates", c)
    }

}
