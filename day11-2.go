package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func (u User) TryToChangeName() {
	u.FirstName = "CHANGED"
	fmt.Println("Inside method, name is :", u.FirstName)
}

func main() {
	user := User{FirstName: "Anda", LastName: "Chore"}
	fmt.Println("Full Name :", user.FullName())

	user.TryToChangeName()
	fmt.Println("Back in main name is :", user.FirstName)
}
