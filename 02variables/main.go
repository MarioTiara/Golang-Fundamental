package main

//Ref ->  https://go.dev/ref/spec#Variables

import "fmt"

// declare var outsite of func can not use non var style //
var globalvar = "the global variable"

func main() {
	var username string = "mario"
	var age int = 24
	fmt.Println(username)
	fmt.Println(age)
	fmt.Printf("variable is of type : %T \n", username)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T \n", smallVal)

	//defaul values and some alies
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "mario.pratama.com"
	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n", website)

	//no var style
	numberOfUser := 20000
	fmt.Println(numberOfUser)
}
