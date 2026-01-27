package main

import "fmt"

func main() {
	fmt.Println()

	fmt.Println("Welcome to functions in Golang")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is", result)

	greeterTwo()

	proResult, myMsg := proAdder(2, 5, 8, 7)

	fmt.Println("Pro result is", proResult)
	fmt.Println("Pro message:", myMsg)

	fmt.Println()
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return  total, "Hello this is pro adder"
}

func greeter() {
	fmt.Println("Hello from Golang")
}

func greeterTwo() {
	fmt.Println("Another method")
}
