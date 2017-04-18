/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Parse Nessus XML file to find all hosts with the specified plugin.

Example usage:

$ go run nessus_plugin.go nessus_file plugin_id
*/

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Nessus struct {
	Reports []Report `xml:"Report"`
}

type Report struct {
	Hosts []Host `xml:"ReportHost"`
}

type Host struct {
	Items      []Item     `xml:"ReportItem"`
	Properties []Property `xml:"HostProperties>tag"`
}

type Property struct {
	Name string `xml:"name,attr"`

	// Caputure the vaule of the <tag> tag.
	Value string `xml:",innerxml"`
}

type Item struct {
	Port     string `xml:"port,attr"`
	Protocol string `xml:"protocol,attr"`
	Id       string `xml:"pluginID,attr"`
}

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run nessus_plugin.go nessus_file plugin_id")
		os.Exit(1)
	}

	// Open the Nessus file
	f, err := os.Open(os.Args[1])
	check(err)
	defer f.Close()

	// Read the data in the file
	data, err := ioutil.ReadAll(f)
	check(err)

	// Unmarshal the data into the nessus structure we built.
	var nessus Nessus
	err = xml.Unmarshal(data, &nessus)
	check(err)

	fmt.Printf("Checking %s for hosts with plugin %s.\n", os.Args[1],
		os.Args[2])
	fmt.Println("=====")

	for _, r := range nessus.Reports {
		for _, h := range r.Hosts {
			var ip string

			for _, p := range h.Properties {
				if p.Name == "host-ip" {
					ip = p.Value
				}
			}

			for _, i := range h.Items {
				if i.Id == os.Args[2] {
					fmt.Printf("%s (%s/%s)\n", ip, i.Port, i.Protocol)
				}
			}
		}
	}
}
