package main 

import (
	"fmt"
)

// we need to only convert string into bytes and bytes into string 

func main() {

	var myBytes []byte
	myBytes = []byte("hello-world")

	fmt.Println(myBytes)

	var myBytesInString string
	myBytesInString = string(myBytes)

	fmt.Println(myBytesInString)
}