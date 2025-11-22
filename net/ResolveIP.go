package main

import (
	"fmt"
	"net"
	"os"
)

func main(){
	if len(os.Args) !=2 {
		fmt.Println("input Error")
		os.Exit(1)
	}
	name := os.Args[1]

// The ResolveIPAddr() will perform a DNS lookup
// on a hostname and return a single IP address 

	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is ", addr.String())
	os.Exit(0)
}