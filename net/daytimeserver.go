package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// The function to create a TCPAddr is ResolveTCPAddr:
// func ResolveTCPAddr(net, addr string) (*TCPAddr, error)
// Where net is one of tcp, tcp4, or tcp6 and the addr is a string composed of a hostname or IP address,
// followed by the port number after a :, such as www.google.com:80 or 127.0.0.1:22
// Another special case is often used for servers, where the host address is zero, so that the TCP
// address is really just the port name, as in :80 for an HTTP server.

//usage : go run daytimeserver.go

func main(){

	servicePort := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servicePort)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// inside a for loop we accept incoming connection
	// a for{} is actually an infinite loop 
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
			// unlike client a server has to run forever
			// so in case there is a error server continues running 
		}
		// client got connected
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func checkError(err error){
	if err != nil {
		fmt.Println(os.Stderr, "fatal error: %s", err.Error())
		os.Exit(1)
	}
}