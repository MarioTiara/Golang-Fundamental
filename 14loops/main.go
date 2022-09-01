package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Firday", "Saturday"}
	fmt.Println(days)

	//common loops
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	//for in range
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }

	for _, day := range days {
		fmt.Printf("value is %v\n", day)
	}
	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 5 {
			goto lco
		} else if rougueValue == 3 {
			rougueValue++
			continue
		}

		fmt.Println("valueis: ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at line 40")

}
