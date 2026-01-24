package main

import "fmt"

func main() {
	fmt.Println()

	fmt.Println("Welcome to Slices in Golang!")

	fruits := []string{"apple", "mango"}

	fmt.Println(fruits)

	fruits = append(fruits, "pear", "banana", "orange")

	fmt.Printf("Fruits: %s of data type: %T \n", fruits, fruits)

	fmt.Println()
}