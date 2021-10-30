package app

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Host string
	Port int
}

func NewConfig() *Config {
	hostName := "localhost"
	envPort := os.Getenv("PORT")

	port, err := strconv.Atoi(envPort)
	if envPort == "" || err != nil {
		port = 8000
	}

	return &Config{
		hostName,
		port,
	}
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
