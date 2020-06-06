package main

import(
	"fmt"
)

type User struct {
	Username string
	Password string
}

func (user *User) login() {
	fmt.Println("Prueba de login")
}

func main() {
	user := &User{"Username", "Password"}
	user.login()
}