package main

import (
	"io"
	"fmt"
	"net/http"
)

const url = "https://jek-app.onrender.com/lessons"

func main() {
	fmt.Println("Golang web request")

	response, err := http.Get(url)

	checkNilError(err)

	fmt.Printf("Response is of type: %T \n", response)

	fmt.Println("Response:", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)

	checkNilError(err)

	content := string(databytes)
	fmt.Println(content)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
