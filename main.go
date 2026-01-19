package main
import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	fmt.Println() // print a new line

	fruits := [4]string{"apple", "orange", "pawpaw", "banana"}
	
	for i := range(len(fruits)) {
		fmt.Println(fruits[i])
	}

}