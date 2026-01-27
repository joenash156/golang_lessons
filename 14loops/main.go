package main

import "fmt"

func main() {
	fmt.Println()

	fmt.Println("Welcome to loops in Golang")

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thurday", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// fmt.Println()

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for i, day := range days {
	// 	fmt.Printf("Index is %v and value is %v \n", i, day)
	// }

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto josh
		}

		fmt.Println(rogueValue)
		rogueValue++
	}

	josh:
		fmt.Println("I am Joshua")

	fmt.Println()
}
