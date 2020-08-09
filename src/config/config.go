package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// TODO: Replace with real server URL when core app goes live
const DEFAULT_SERVER_URL = "http://localhost:2000"

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
