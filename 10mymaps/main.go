package main

import (
	"fmt"
)

func main() {
	fmt.Println()

	fmt.Println("Welcome to the lesson on Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS is short for: ", languages["JS"])

	delete(languages, "RB")

	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For kay %v, value is %v \n", key, value)
	}

	fmt.Println()
}
