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

func (argPerson Person) details() string {
	// return "a detailed description"
	var description string
	description = argPerson.sName.firstName +
				  argPerson.sName.lastName + "-"

	for _, eachEmail := range argPerson.sEmail {
		var emailPair = eachEmail.addr + "-" + eachEmail.kind
		description += emailPair
	}

	return description
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

	fmt.Println(aPerson.details())

}