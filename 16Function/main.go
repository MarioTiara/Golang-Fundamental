package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Function in Golang")
	VoidFunc()
	VoidFunc2("mario", 24)
	identity := returnFunc("mario", 24)
	fmt.Println(identity)
}

func VoidFunc() {
	fmt.Println("this void without input")
}

func VoidFunc2(name string, age int) {
	fmt.Printf("Name: %v Age: %v", name, age)
}

func returnFunc(name string, age int) string {
	Identity := "Name: " + name + " Age: " + strconv.Itoa(age)
	return Identity

}
