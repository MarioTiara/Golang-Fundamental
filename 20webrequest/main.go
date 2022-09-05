package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://pkg.go.dev/net/http"

func main() {
	fmt.Println("LCO web request")
	reponse, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", reponse)
	fmt.Println(reponse.Status)
	databytes, err := ioutil.ReadAll(reponse.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
	reponse.Body.Close()
}
