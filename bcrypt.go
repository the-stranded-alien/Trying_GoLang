package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(string(bs))

	loginPword1 := "password123"
	er := bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if er != nil {
		fmt.Println("You can't login")
		return
	}
	fmt.Println("You are logged in")
}
