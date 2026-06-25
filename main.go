package main

import (
	"fmt"
	"log"
	"os"
)

type IPInfo struct {
	Bogon    bool   `json:"bogon"`
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func main() {
	if len(os.Args) < 2 {
		log.Println("usage: go run . <ip_address>")
		return
	}

	data, err := LookupIP(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IP       : %s\n", data.IP)
	fmt.Printf("Country  : %s\n", data.Country)
	fmt.Printf("Region   : %s\n", data.Region)
	fmt.Printf("City     : %s\n", data.City)
	fmt.Printf("Org      : %s\n", data.Org)
	fmt.Printf("Timezone : %s\n", data.Timezone)
	fmt.Printf("Location : %s\n", data.Loc)
}
