package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println()

	fmt.Println("Welcome to files in Golang")

	content := "This needs to go in a file. Yes - josh.com"

	file, err := os.Create("./myfile1.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is", length)

	defer file.Close()
	// readFile("./myfile1.txt")

	fmt.Println()
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

// func readFile(filename string) {
// 	databyte, err := ioutil.ReadFile(filename)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Text data inside the file \n", databyte)
// }
