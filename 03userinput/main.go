package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println() // print a new line

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	name, err := reader.ReadString('\n')
	fmt.Println(err)
	fmt.Printf("Thank you, %s \n", name)


	//fmt.Println() // print a new line
}
