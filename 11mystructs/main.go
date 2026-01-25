package main

import "fmt"

func main() {
	fmt.Println()

	fmt.Println("Structs in golang")

	// no inheritance in golang; no super or parent class concepts

	josh := User{
		"Joshua",
		"joenash@gmail.com",
		true,
		16,
	}

	fmt.Println(josh)
	fmt.Printf("Josh's details are: %+v \n", josh)

	fmt.Printf("Name is: %v and Email is: %v \n", josh.Name, josh.Email)

	var user1 User
	var name = user1.Name

	name = "Kwame"

	fmt.Println(name)

	fmt.Println()
}

// a struct (structure) for the user
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
