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

	fmt.Println("WELCOME TO MY SIMPLE CALCULATOR")
	fmt.Println("===============================")
	fmt.Println() // print a new line

	options := []string{"Addition (+)", "Subtraction (-)", "Division (/)", "Multiplication (x)"}

	for i, option := range options {
		fmt.Printf("%d. %s \n", i+1, option)
	}

	fmt.Println() // print a new line

	// take users option
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose an operation: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured: ", err)
		return
	}

	option, err := strconv.ParseInt(strings.TrimSpace(input), 10, 8)

	if err != nil {
		fmt.Println("An error occured: ", err)
		return
	}

	fmt.Println() // print a new line

	fmt.Println("You chose to do : ", options[option - 1])

	fmt.Println() // print a new line

	var num1, num2, answer float64

	for i := 1; i < 3; i++ {
		fmt.Printf("Enter num%d number: ", i)
		input, _ = reader.ReadString('\n')

		if i == 1 {
			num1, _ = strconv.ParseFloat(strings.TrimSpace(input), 64)

		} else {
			num2, _ = strconv.ParseFloat(strings.TrimSpace(input), 64)
		}
	}

	fmt.Println() // print a new line

	switch option {
		case 1:
			answer = addtion(num1, num2)
			fmt.Printf("%.2f + %.2f = %.2f \n", num1, num2, answer)
		
		case 2:
			answer = subtraction(num1, num2)
			fmt.Printf("%.2f - %.2f = %.2f \n", num1, num2, answer)

		case 3:
			answer = division(num1, num2)
			fmt.Printf("%.2f / %.2f = %.2f \n", num1, num2, answer)

		case 4:
			answer = multiplication(num1, num2)
			fmt.Printf("%.2f x %.2f = %.2f \n", num1, num2, answer)
	}
	

	fmt.Println() // print a new line
}
