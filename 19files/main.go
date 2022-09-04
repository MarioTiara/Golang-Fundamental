package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "this is content to write in file"

	file, err := os.Create("./19files/myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Lenght is: ", length)
	defer file.Close()
	readFle("./19files/myfile.txt")
}

func readFle(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("text data is:  ", string(databyte))
}
