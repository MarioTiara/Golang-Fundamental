package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input:"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	//comma ok || err ok
	input, _ := reader.ReadString('\n') //-> if input is rigt then "input" will be filled, other wise "-" will be filled by error message
	fmt.Println("thanks for rating, ", input)
	fmt.Printf("type of this rating is %T", input)
}
