package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println()

	fmt.Println("Welcome to Slices in Golang!")

	fruits := []string{"apple", "mango"}

	fmt.Println(fruits)

	fruits = append(fruits, "pear", "banana", "orange")

	fmt.Printf("Fruits: %s of data type: %T \n", fruits, fruits)

	fmt.Printf("Fruits slice has a length: %d and capacity of: %d \n", len(fruits), cap(fruits))

	fmt.Println(sort.StringsAreSorted(fruits))

	sort.Strings(fruits)

	fmt.Println(sort.StringsAreSorted(fruits))

	// how to remove a value from a slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby", "golang"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index + 1:]...)
	fmt.Println(courses)

	fmt.Println()
}