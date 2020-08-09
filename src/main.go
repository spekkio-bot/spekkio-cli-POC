package main

import (
	"fmt"

	"github.com/spekkio-bot/spekkio-cli/src/config"
)

func main() {
	config := &config.Config{}
	config.Initialize()
	fmt.Println(config.ServerUrl)
}
