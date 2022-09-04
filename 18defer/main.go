package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("two")
	fmt.Println("Hello")
	myDefer()
}

//hello, 43210, two, one, world

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
