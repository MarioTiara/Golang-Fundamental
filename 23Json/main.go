package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `josn:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json in Go")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"Golang bootcamp", 299, "mario.pratama.com", "abcd123", []string{"Programming", "go"}},
		{"Dotnet bootcamp", 299, "mario.pratama.com", "abcd123", []string{"web-dev", "C#"}},
		{"C++ bootcamp", 299, "mario.pratama.com", "abcd123", nil},
	}

	//finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
