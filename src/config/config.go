package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	DEFAULT_SERVER_URL = "https://5ila6fw37k.execute-api.us-west-1.amazonaws.com/api"
)

type Config struct {
	ServerUrl string
}

func (c *Config) Initialize() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("no .env file found, using defaults\n")
		c.UseDefault()
	} else {
		c.ServerUrl = os.Getenv("SERVER_URL")
	}
}

func (c *Config) UseDefault() {
	c.ServerUrl = DEFAULT_SERVER_URL
}
