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

    fmt.Printf("customer: %#v \n", customer)
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

    // NOTICE: in this case where return value is an interface and -
    // we're returning a nil pointer (as we didn't initialize MultiError)
    // in this case its type is (*main.MultiError)(nil), here 1st paren is the error (interface) -
    // which is not nil but it points to nil.

    // so the solution can be, do not return nil interface.
    // return a value (of interface) when that is not nil, or the value nil

    if m != nil {
        return m
    }

    return nil
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
