package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println()

	fmt.Println("Welcome to web server in Golang")
	fmt.Println()

	PerformGetRequest()

	fmt.Println()
}

func PerformGetRequest() {
	const myurl = "https://jek-app.onrender.com/lessons"

	response, err := http.Get(myurl)

	checkNilError(err)

	// close response
	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	// dfmt.Println("Content length is:", response.ContentLength)

	var responseString strings.Builder

	content, err := io.ReadAll(response.Body)
	checkNilError(err)

	byteCount, _ := responseString.Write(content)
	checkNilError(err)

	fmt.Println("Byte count is:", byteCount)

	fmt.Println(string(content))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}