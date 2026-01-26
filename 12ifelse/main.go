package main

import "fmt"

func main() {
	fmt.Println()

	fmt.Println("If else in Golang")

	loginCount := 25
	var result string

	if loginCount < 10 {
		result = "Regular user"

	} else if loginCount > 10 {
			result = "Watch out"

	} else {
			result = "Exactly 10 login ocunt"
	}

	fmt.Println("Result: ", result)

	if num := 3; num < 10 {
		fmt.Println("Num is less that 10")
	} else {
			fmt.Println("Number is NOT less than 10")
	}

	fmt.Println()
}
