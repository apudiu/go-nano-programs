package main

import (
	"fmt"
	"net/http"
)

type Config struct {
	Port *int
}

func NewServer(addr string, cfg Config) (*http.Server, error) {
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", addr, *cfg.Port),
		Handler: nil,
	}, nil
}

// problems with this approach
// 1. it’s not handy for clients to provide an integer pointer. Clients have to create a variable and then pass a pointer
// 2. client using this with the default configuration will need to pass an empty struct
