package main

import "fmt"

func main() {
	fmt.Println() // print a new line

	fmt.Println("Welcome to a class on Pointers")

	// var myptr *int
	// fmt.Println("Value of pointer is: ", myptr)

	myNumber := 23

	var myNumberPtr = &myNumber
	fmt.Println("Value of actual pointer is: ", myNumberPtr)
	fmt.Println("Value of actual pointer is: ", *myNumberPtr)

	fmt.Println() // print a new line
}
