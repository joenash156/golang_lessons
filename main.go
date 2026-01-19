package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println() // print a new line

	// fruits := [4]string{"apple", "orange", "pawpaw", "banana"}

	// for i := range(len(fruits)) {
	// 	fmt.Println(fruits[i])
	// }

	// an array to store student score (dammy data)
	studentsScores := [5]int{78, 65, 90, 82, 70}

	// convert to slice
	scores := studentsScores[0:]
	total := 0
	var average float32 = 0

	// print the scores
	fmt.Print("Scores: ")

	for _, score := range scores {
		fmt.Printf("%v ", score)

		// add each score to total
		total += score
	}

	fmt.Println() // print a new line

	// calculate average
	average = float32(total) / float32(len(scores))

	// print the total and avarage
	fmt.Printf("Total: %d \n", total)
	fmt.Printf("Average: %.2f \n", average)

	passed := 0
	failed := 0

	for _, score := range scores {
		if score >= 50 {
			passed++
		} else {
			failed++
		}
	}

	// print number of passed and failed
	fmt.Printf("Passed: %d \n", passed)
	fmt.Printf("Failed: %d \n", failed)

	// slice to store above average scores
	aboveAverage := []int{}

	for _, score := range scores {
		if float32(score) > average {
			aboveAverage = append(aboveAverage, score)
		}
	}

	// print above average slice
	fmt.Print("Above Average Scores: ")
	for _, score := range aboveAverage {
		fmt.Printf("%d ", score)
	}

}
