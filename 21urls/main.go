package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://local.dev:300/learn?coursename=golang&paymentid=ghj456ghb"

func main() {
	fmt.Println("url in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of query params are : %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "local.dev",
		Path:    "/tutcss",
		RawPath: "user=mario",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
