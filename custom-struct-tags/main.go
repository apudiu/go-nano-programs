package main

import (
	"customstructtags/validator"
	"fmt"
	"log"
)

func main() {
	user := User{
		Name:  "AB",
		Email: "ab@c.d",
	}

	if err := validator.Validate(user); err != nil {
		fmt.Printf("errors: %v \n", err)
		return
	}

	log.Println("No validation errors 😀")
}

type User struct {
	Name  string `validate:"required,min=2,max=10"`
	Email string `validate:"required,email"`
}
