package main

import (
	"fmt"
	"net"
	"os"
)

// we create a function to check fatal error and exit

func checkError(err error){
	if err != nil {
		fmt.Println(os.Stderr, "fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// we create a function which reads the received bytes and replies the same 

func handleClient(conn net.Conn){
	var buf [512]byte
	// the following infinite loop is a imp pattern
	// possibly it breaks when no more bytes is left 
	for {
			n, err := conn.Read(buf[0:])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(buf[0:]))
			_, err2 := conn.Write(buf[0:n])
			if err2 != nil {
				fmt.Println(err2)
				return
			}
	}
}

// usage : go run simpleEchoServer.go
// another terminal $ telnet 127.0.0.1 1201
//                  $ hello 
//                  Enter

func main(){

	// we resolve our service and get the tcpAddr 
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	// now we create a listener from the obtained tcpAddr
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err) //look for critical error at this stage

	// in a for infinite loop we listen and accept incoming connection
	for{
		conn, err := listener.Accept()
		if err != nil {
			continue // there is some error so continue listening for next connection
		}
		handleClient(conn)
		conn.Close() // we are finished 
	}
}