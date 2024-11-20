package config

import "os"

func init() {
	if os.Getenv("PORT") != "" {
		Conf.Port = os.Getenv("PORT")
	}
}

type Config struct {
	Https bool
	Host  string
	Port  string
}

func (c *Config) GetUrl() string {
	u := "http://"
	if c.Https {
		u = "https://"
	}
	u += c.Host + ":" + c.Port
	return u
}

func (c *Config) GetWsUrl() string {
	u := "ws://"
	if c.Https {
		u = "wss://"
	}
	u += c.Host + ":" + c.Port + "/ws"
	return u
}

var Conf = &Config{
	Https: false,
	Host:  "localhost",
	Port:  "8000",
}
