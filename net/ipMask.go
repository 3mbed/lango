package main 

import (
	"fmt"
	"net"
	"os"
	"strconv"
)
// A IP address is typically divided into the components
// network address + subnet + device portion 

// IPMask type is also an array of bytes 

func main(){

	if len(os.Args) != 4 {
		fmt.Println(os.Stderr, "Usage: %s dotted-ip-addr ones bits\n", os.Args[0])
		os.Exit(1)
	}
	// := is for declaration + assignment
	//  = is for assignment only 

	dotAddr := os.Args[1]
	ones, _ := strconv.Atoi(os.Args[2])
	bits, _ := strconv.Atoi(os.Args[3])

	// args printing 
	fmt.Println(dotAddr)
	fmt.Println(ones)
	fmt.Println(bits)

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid Address")
		os.Exit(1)
	}

	mask := net.CIDRMask(ones, bits)
	network := addr.Mask(mask) 

	fmt.Println("Address is ", addr.String(),
               "\nMask length is ", bits,
               "\nLeading ones count is ", ones,
               "\nMask is (hex) ", mask.String(),
               "\nNetwork is ", network.String())
    os.Exit(0)
}