package app

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port int
	Addr string
}

func NewConfig() *Config {
	hostName := "localhost"
	envPort := os.Getenv("PORT")

	port, err := strconv.Atoi(envPort)
	if envPort == "" || err != nil {
		port = 8000
	}

	addr := fmt.Sprintf("%s:%d", hostName, port)

	return &Config{
		port,
		addr,
	}
}
