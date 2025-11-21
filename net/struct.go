package main

import(
	"fmt"
)

type Email struct {
	addr string
	kind string
}

type Name struct {
	firstName string
	lastName string
}

type Person struct{
	sName  Name
	sEmail []Email
}

func main(){
	fmt.Println("using struct for email")

	var aName 	Name
	var aEmail  Email
	var bEmail  Email
	var arrEmail []Email
	var aPerson Person

	fmt.Println(aName)
	fmt.Println(aEmail)
	fmt.Println(bEmail)
	fmt.Println(arrEmail)
	fmt.Println(aPerson)

	aName    = Name{"Ayman", "Ahmed"}
	aEmail   = Email{"xyz@gmail.com", "Home"}
	bEmail   = Email{"abc@gmail.com", "work"}

	arrEmail = []Email{
		aEmail,
		bEmail,
	}

	aPerson = Person{ 
		aName, 
		arrEmail, 
	}

	fmt.Println(aName)
	fmt.Println(aEmail)
	fmt.Println(bEmail)
	fmt.Println(arrEmail)
	fmt.Println(aPerson)
}