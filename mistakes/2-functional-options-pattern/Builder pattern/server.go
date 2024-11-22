package main

import (
	"fmt"
	"net/http"
)

type Config struct {
	Port int
}

type ConfigBuilder struct {
	port *int
}

func (c *ConfigBuilder) Port(p int) *ConfigBuilder {
	c.port = &p
	return c
}

func (c *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}
	if c.port == nil { // when port is not provided use default
		defaultPort := 3000
		cfg.Port = defaultPort
	} else {
		if *c.port == 0 { // when port is 0 use random port
			randomPort := 4000
			cfg.Port = randomPort
		} else if *c.port < 0 { // validate passed port
			return cfg, fmt.Errorf("port %d must be positive", *c.port)
		} else { // when provided port is okay, use that
			cfg.Port = *c.port
		}
	}

	return cfg, nil
}

func NewServer(addr string, cfg Config) (*http.Server, error) {
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", addr, cfg.Port),
		Handler: nil,
	}, nil
}

// problems with this approach
// 1. need to pass a config struct that can be empty if a client wants to use the default configuration
