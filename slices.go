package main

import (
	"fmt"
	// "slices"
)

func main(){
	
	var s[]string
	// s := "hello"
	fmt.Println(s)

	fmt.Println("uninit:", s, s==nil, len(s) == 0)

	// now we create a slice
	s = make([]string, 3 )
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)
	fmt.Println(s[2])
	fmt.Println(len(s))
	
	s = append(s, "d")
	s = append(s, "dr")
	s = append(s, "drsss")
	fmt.Println(s)
	fmt.Println(len(s))

	c:= make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	fmt.Println(c[0:2])
	fmt.Println(s[3:5])

	str1 := "hello"
	str2 := "hello"

	if(str1 == str2){
		fmt.Println(str1 == str2)
	}
}