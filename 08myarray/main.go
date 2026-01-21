package main

import "fmt"

func main() {
	fmt.Println() // print a new line

	fmt.Println("Welcome to array in golang")

	var fruits [4]string

	fruits[0] = "apple"
	fruits[1] = "tomato"
	fruits[3] = "peach"

	fmt.Println(fruits)

	var vegetables = [5]string{"potato", "beans", "mushroom"}

	fmt.Println("Vegetable list is: ", vegetables)

	fmt.Println() // print a new line
}
