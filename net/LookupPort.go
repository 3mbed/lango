// services live on a host machine
// there can be many services in a host machine
// port number distinguishes services on a host
// port number range could be 1 and 65536 

// for example http uses port 80, 8000, 8080 , 8088 

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Input Error")
		os.Exit(1)
	}
	networkType := os.Args[1]
	service := os.Args[2]

	port, err := net.LookupPort(networkType, service)

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(2)
	}

	fmt.Println("service port: ", port)
	os.Exit(0)
}

// go run LookupPort tcp telnet
// go run LookupPort tcp ssh
// go run LookupPort tcp ftp
// go run LookupPort tcp http