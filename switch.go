package main 

import (
	"fmt"
	"time"
)

func main(){

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's afternoon")
	}

	// now we create a variable and pass a function to it
	whatAmI := func( i interface{}){
		switch tyPe := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Println(" unknown type %T", tyPe)
		}
	} 

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}