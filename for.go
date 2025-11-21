package main

import "fmt"

func main() {
	i := 1
	for i < 3{
		fmt.Println(i)
		i = i+1
	}
	fmt.Println("---------------")
	for j:=0; j<3; j++ {
		fmt.Println(j)
	}
	fmt.Println("---------------")
	fmt.Println(i)
	for i := range 3 {
		fmt.Println("Range", i)
	}
	for {
        fmt.Println("loop")
        break
    }
   	for n := range 6 {
        if n%2 != 0 {
            continue
        }
        fmt.Println(n)
    }
}