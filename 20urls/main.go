package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://jek-app.onrender.com:3000/lessons?section=one&lesson_id=ls001"

func main() {
	fmt.Println()

	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myUrl)

	// parsing
	result, err := url.Parse(myUrl)
	checkNilError(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T \n", qparams)

	fmt.Println(qparams["section"])

	for key, val := range qparams {
		fmt.Printf("%v : %v \n", key, val)
	}

	// construct custom url
	partsOfUrl := &url.URL {
		Scheme: "https",
		Host: "josh.dev",
		Path: "/tutorial",
		RawPath: "user=josh",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println("Another url:", anotherUrl)

	fmt.Println()
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
