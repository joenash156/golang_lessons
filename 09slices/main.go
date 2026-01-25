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

	fmt.Println()
}