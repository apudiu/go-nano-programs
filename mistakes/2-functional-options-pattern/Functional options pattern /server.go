package main

import (
	"fmt"
	"net/http"
)

type options struct {
	port *int
}

type Option func(*options) error

func WithPort(port int) Option {
	return func(o *options) error {
		if port < 0 || port > 65535 {
			return fmt.Errorf("port should be between %d & %d", 0, 65535)
		}

		o.port = &port
		return nil
	}
}

func NewServer(addr string, opts ...Option) (*http.Server, error) {

	cfg := options{}
	for _, opt := range opts {
		if err := opt(&cfg); err != nil {
			return nil, err
		}
	}

	// options validation
	if cfg.port == nil { // when port is not provided use default
		defaultPort := 3000
		cfg.port = &defaultPort
	} else {
		if *cfg.port == 0 { // when port is 0 use random port
			randomPort := 3000
			cfg.port = &randomPort
		}
	}

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", addr, *cfg.port),
		Handler: nil,
	}, nil
}

// problems with this approach
// none
