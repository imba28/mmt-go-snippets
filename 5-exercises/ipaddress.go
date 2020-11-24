package main

import "fmt"

type IPAddr [4]byte

// TODO: add IPAddrInterface

// TODO: add ToString method for IPAddr

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
	}
}
