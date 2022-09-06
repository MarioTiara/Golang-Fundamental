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
	//EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename":"Golang Bootcamp",
			"Price":299,
			"webise":"mario.pratama.com",
			"tags":["web-dev", "js"]
		}
	`)

	var lcoCourses course

	checkvalid := json.Valid(jsonDataFromWeb)
	if checkvalid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("Json was not valid")
	}

	// some case where you just want to addd data key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("ker is %v and value is %v and type is : %T\n", k, v, v)
	}

}
