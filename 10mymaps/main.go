package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"strings"
)

func main() {
	fmt.Println()

	fmt.Println("Welcome to the lesson on Maps in Golang")

	// languages := make(map[string]string)

	// languages["JS"] = "JavaScript"
	// languages["RB"] = "Ruby"
	// languages["PY"] = "Python"
	// languages["GO"] = "Golang"

	// fmt.Println("List of all languages: ", languages)
	// fmt.Println("JS is short for: ", languages["JS"])

	// delete(languages, "RB")

	// fmt.Println("List of all languages: ", languages)

	// // loops are interesting in golang
	// for key, value := range languages {
	// 	fmt.Printf("For kay %v, value is %v \n", key, value)
	// }

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Which coutry do you want the capital of: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error occured getting capital city of %s \n", input)
		return
	}

	input = strings.TrimSpace(input)

	capital := countryCapital(input)

	fmt.Printf("The capital city of %s is %s \n", input, capital)

	// fmt.Println(capital)

	fmt.Println()
}
