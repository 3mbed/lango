package main 

import (
	"fmt"
)

func sums( a_val int) int{
	return a_val + 1
}

func sumss( a_val ...int) int{
	return len(a_val)
}

func sumss_total( a_val ...int) int{
		
	sums := 0
	for val := range a_val {
		// fmt.Println(val)
		sums = sums + val
	}

	return sums
	// return len(a_val)
}

func main(){
	// fmt.Println(sums(5))

	arr := []int{1,2,3,4}
	
	// fmt.Println(arr)
	// fmt.Println(sumss(arr...))
	fmt.Println(sumss_total(arr...))
}