package main

import (
	"fmt"
)

func main() {

	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of all language: ", languages)
	fmt.Println("Js Short For: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all language: ", languages)

	//loop are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, values is %v\n", key, value)
	}
}
