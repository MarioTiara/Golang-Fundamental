package main

import "fmt"

func main() {
	mario := User{"Mario", "hitesh@go.dev", true, 16}
	fmt.Println(mario)
	mario.GetStatus()
	mario.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Ïs user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
