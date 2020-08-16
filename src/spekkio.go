package main

import (
	"os"

	"github.com/spekkio-bot/spekkio-cli/src/cli"
	"github.com/spekkio-bot/spekkio-cli/src/config"
)

func main() {
	cli := &cli.CliApp{
		Config: &config.Config{},
		Args:   os.Args,
	}
	cli.Initialize()
	cli.Run()
}
