package main

import (
	"fmt"
)

type User struct {
	login string
	password string
}

func (u * User) setLogin(login string) {
	*&u.login = login
}

func main() {
	user1 := User{
		"nikita",
		"t0slrl",
	}

	user1.setLogin("satoshi")
	fmt.Println(user1.login)

}