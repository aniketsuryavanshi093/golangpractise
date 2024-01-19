package main

import "fmt"

func main() {
	aniket := User{"aniket", "test@gmail.com", true, 12}
	fmt.Println("name and email is:", aniket.Name, aniket.email)
	aniket.GetStatus()
	aniket.newMail()
	fmt.Println("name and email after:", aniket.Name, aniket.email)

}

type User struct {
	Name   string
	email  string
	status bool
	age    int
}

// methods
func (u User) GetStatus() {
	fmt.Println("is user active:", u.status)
}

func (u User) newMail() {
	u.email = "test@go.dev"
	fmt.Println("email of this user is :", u.email)
}
