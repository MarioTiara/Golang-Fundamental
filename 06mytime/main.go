package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presenttime := time.Now()
	fmt.Println(presenttime)

	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.Local)
	fmt.Println(createDate)
}
