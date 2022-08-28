package main

import "fmt"

func main() {
	fmt.Println(" Pointer in Golang")

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("value of actual pointer is, ", ptr)
	fmt.Println("value of depointer is, ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value is: ", myNumber)
}
