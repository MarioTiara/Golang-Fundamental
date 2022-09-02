package main

import "fmt"

func main() {

	for i := 0; i <= 5; i++ {
		for j := 0; j <= 3; j++ {
			if i == j {
				//break
				continue
			}
			fmt.Printf("i:%v j:%v \n", i, j)
		}
	}

}
