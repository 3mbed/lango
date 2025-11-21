//ipv4 address is a 32 bit Integer 
//it addresses down to a network interface card on a single device

// IP address is composed of two parts
// (1) The address of the network in which device resides
// (2) The address of the device within the network 

// class A network : First Byte => identifies the network
// class B network : First 2 Bytes => identifies the network 
// Class C network : First 3 bytes => identifies the network

// to find the network of a device
// bitwise AND its IP address with the network mask 

// to find the device address within the subnet 
// create one's complement of the network mask 
// then bitwise AND it with the IP address 

// for 16 bit network address , netmask is 255.255.0.0
// for 24 bit network address , netmask is 255.255.255.0
// for 23 bit network address , netmask is 255.255.254.0
// for 22 bit network address , netmask is 255.255.252.0

// the type IP is defines as an array of Bytes 

package main 

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2{
		fmt.Println(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
	}
	name := os.Args[1]
	fmt.Println(name)

	addr := net.ParseIP(name)
	fmt.Println(addr)

	if addr == nil {
		fmt.Println("Invalid address")
	}else{
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}

