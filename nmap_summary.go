/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Parse an Nmap XML file to get a summary of the data.

Example usage:

$ go run nmap_summary.go nmap_file

I used the example here: https://gist.github.com/kwmt/6135123#file-parsetvdb-go
as the starting point for this code.

More details about parsing XML can be found here: https://golang.org/pkg/encoding/xml
*/

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

/*
Go attempts to unmarshal the XML data into the structs you provide. When the
struct names and variable names match the XML tags, Go can usually figure out
how to unmarshal the data itself. Sometimes though you have to provide hints.
The hints are the `xml:...` bits of code after each variable name. The best I
can tell you will always need to provide hints when the data you want is an
attribute. You will also need to provide hints when the variable name does not
match the XML tag.

You will notice that all the variables inside the structs are capitalized.
This is necessary because of the way Go unmarshals the data in the background.
I do not understand the details and have no desire to understand them at this
time. I'm just doing what I know works. :)

Also note that each struct does not contain every piece of data available. Go
will populate the data you ask for and ignore the rest. This means you can
modify this program to capture more data like the script output stored in the
<script> tag.
*/

type Nmap struct {
	// The root element <nmaprun> has a version and args attribute. We will
	// capture those.
	Version string `xml:"version,attr"`
	Args    string `xml:"args,attr"`

	// The Nmap file has a number of <host> tags, each one will go in its own
	// struct.
	Hosts []Host `xml:"host"`
}

// Take notice that because Go is attempting to populate the Host struct with
// data from the <host> tag, our hints assume we are inside a <host> tag. We
// do not have to reference host>address in the hint.
type Host struct {
	// Each <host> has multiple address tags, capture each one in its own struct
	Addresses []Address `xml:"address"`

	// Each of the <port> tags is nested in the <ports> tag. We need to tell
	// Go how about this nesting.
	Ports []Port `xml:"ports>port"`

	// Each host has a <status> tag. Put it in a struct.
	Status State `xml:"status"`
}

type Address struct {
	// Each <address> tag has an addr and addrtype attribute.
	Address     string `xml:"addr,attr"`
	AddressType string `xml:"addrtype,attr"`
}

type Port struct {
	// Each <port> tag has a portid and protocol attribute.
	Port     string `xml:"portid,attr"`
	Protocol string `xml:"protocol,attr"`

	// Each <port> tag has a <state> tag and a <service> tag.
	State   State   `xml:"state"`
	Service Service `xml:"service"`
}

type State struct {
	// Both the <state> and <status> tags use the same format so we can use
	// The same struct for both.
	State  string `xml:"state,attr"`
	Reason string `xml:"reason,attr"`
}

type Service struct {
	// Grab the service details from the attributes of the <service> tag.
	Name    string `xml:"name,attr"`
	Product string `xml:"product,attr"`
	Version string `xml:"version,attr"`
}

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run nmap_summary.go nmap_file")
		os.Exit(1)
	}

	// Open the Nmap file
	f, err := os.Open(os.Args[1])
	check(err)
	defer f.Close()

	// Read the data in the file
	nmap_data, err := ioutil.ReadAll(f)
	check(err)

	// Unmarshal the data into the nmap structure we built.
	var nmap Nmap
	err = xml.Unmarshal(nmap_data, &nmap)
	check(err)

	// Once the data is unmarshaled into its various structs we can use
	// standard struct notation to get to the data.
	fmt.Printf("Summary of %s\n", os.Args[1])
	fmt.Println("=====")
	fmt.Printf("Arguments: %s\n", nmap.Args[5:])
	fmt.Printf("Version: %s\n", nmap.Version)

	for _, h := range nmap.Hosts {
		for _, a := range h.Addresses {
			if a.AddressType == "ipv4" || a.AddressType == "ipv6" {
				fmt.Printf("%s\n", a.Address)
			}
		}

		fmt.Println("-----")
		fmt.Printf("Status: %s\n", h.Status.State)

		for _, p := range h.Ports {
			fmt.Printf("%s/%s - %s %s %s\n", p.Port, p.Protocol,
				p.Service.Name, p.Service.Product, p.Service.Version)
		}

		fmt.Println()
	}
}
