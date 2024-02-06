package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	var str_list []string
	for _, number := range ip {
		str_list = append(str_list, fmt.Sprintf("%d", number))	
	}
	return strings.Join(str_list, ".") 

	// return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
	
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
