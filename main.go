package main
import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	println() // print a new line

	name := "Joshua"
	var age = 24

	var pers1 Person
	
	pers1.name = "Kwesi"
	pers1.age = 23

	fmt.Println(pers1.name, pers1.age)

	fmt.Printf("My name is %s and I am %d years of age \n", name, age)

	println() // print a new line
}