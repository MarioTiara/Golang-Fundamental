package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to get request with go")
	//PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"
	//fake json pyload
	requestBody := strings.NewReader(`
	{
		"coursename":"Let's go with golang",
		"price": 0,
		"platform":"learnCodeOnline.in"
	}`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content lenght:", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	// fmt.Println(content)
	// fmt.Println(string(content))

	fmt.Println("Bytecount is: ", byteCount)
	fmt.Println(responseString.String())
}
