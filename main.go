package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Premium bool
}

func main() {
	user := User{
		Name:    "Fedor",
		Age:     18,
		Premium: true,
	}

	fmt.Println(user.Name)
}
