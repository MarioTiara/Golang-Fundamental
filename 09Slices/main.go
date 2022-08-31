package main

import "fmt"

func main() {
	fruits := []string{"apple", "grape", "banana", "melon"}
	newfruits := fruits[:3]

	newfruits = append(newfruits, "mango", "pear")

	fmt.Println(newfruits)

	//how to remove a value from a slices

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
