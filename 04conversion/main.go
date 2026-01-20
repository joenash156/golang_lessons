package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println() // print a new line

	fmt.Println("Welcome to our Pizza App")
	fmt.Println() // print a new line

	fmt.Print("Please rate our pizza between 1 and 10: ")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')

	// fmt.Println("An error occured: ", err)

	fmt.Printf("Thank you for rating: %s \n", rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println("Error occured!: ", err)
		
	} else {
		fmt.Println("Added 1 to your rating: ", numRating + 1)
	}

	fmt.Println() // print a new line
}
