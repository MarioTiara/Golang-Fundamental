package main

import "fmt"

func main() {
	fmt.Println("If else in golang")
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "something else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Numbers is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Nu is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

}
