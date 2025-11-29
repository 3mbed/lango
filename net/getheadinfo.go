package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

//usage
//go run getheadinfo.go www.google.com:80

func checkError(err error){
	if err != nil {
		fmt.Println(os.Stderr, "fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main(){

	fmt.Println(os.Args[0])

	if len(os.Args) !=2 {
		fmt.Println("Error")
		os.Exit(1)
	}
	service := os.Args[1]

	targetTcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, targetTcpAddr)
	checkError(err)

	var msg2send []byte
	msg2send = []byte("HEAD / HTTP/1.0\r\n\r\n")	
	_, err = conn.Write(msg2send)
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}