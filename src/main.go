package main

import (
	"github.com/spekkio-bot/spekkio-cli/src/cli"
	"github.com/spekkio-bot/spekkio-cli/src/config"
)

func main() {
	cli := &cli.CliApp{
		Config: &config.Config{},
	}
	cli.Initialize()
}
