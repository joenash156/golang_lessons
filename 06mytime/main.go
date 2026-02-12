package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println() // print a new line

	fmt.Println("Welcome to Time Study of Golang!")
	fmt.Println() // print a new line

	presentTime := time.Now()

	fmt.Println(presentTime)

	// change time format
	fmt.Println(presentTime.Format("01/02/2006 15:04:05 Monday"))

	createdDate := time.Date(2002, time.April, 22, 23, 30, 0, 0, time.UTC)
	fmt.Println("Date: ", createdDate.Format("01-02-2006 15:04:05 Monday"))

	fmt.Println() // print a new line
}
