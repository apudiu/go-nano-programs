package main

import (
	"customstructtags/validator"
	"fmt"
)

func main() {
	user := User{
		Name:  "A",
		Email: "abcd",
	}

	if err := validator.Validate(user); err != nil {
		fmt.Println("errors:", err)
	}
}

type User struct {
	Name  string `validate:"required,min=2,max=10"`
	Email string `validate:"required,email"`
}
