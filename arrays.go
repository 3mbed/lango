package main

import "fmt"


func check_values(arr [5]int){
	fmt.Println(arr)
} 

func example_2D_arr(){
	var twoD[2][3] int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println(twoD)
}

func example_2D_arr_2(){
	var twoD[2][3]int 

	twoD =[2][3]int {
		{1,2,3},
		{4,5,6},
	}
	fmt.Println(twoD)
}

func example_arr_as_ptr(){
	xp := new([12]int)
	fmt.Println(xp)

	for i:=0; i<len(xp); i++{
		xp[i] = 5
	}

	fmt.Println(xp)
}

func main(){

	var a[5] int
	// fmt.Println(a)
	check_values(a)

	example_2D_arr()
	example_2D_arr_2()

    // ---------------------------------------
	whatAmI := func( i interface{}){
		switch tyPe := i.(type) {
		default:
			fmt.Println(" unknown type %T", tyPe)
		}
	}

	whatAmI(a)
	fmt.Println("-----------------------")

	example_arr_as_ptr()
}