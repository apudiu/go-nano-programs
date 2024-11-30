package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	customer := Customer{Age: 33, Name: "John"}
	if err := customer.Validate(); err != nil {
		log.Fatalf("customer is invalid: %#v", err)
	}

	fmt.Printf("%#v \n", customer)
}

type Customer struct {
	Age  int
	Name string
}

func (c Customer) Validate() error {
	var m *MultiError

	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}

	// NOTICE: this is not nil but a nil pointer to a interface
	return m

}

// ------------

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}
