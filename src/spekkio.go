package main

import (
	"os"

	"github.com/spekkio-bot/spekkio-cli/src/cli"
	"github.com/spekkio-bot/spekkio-cli/src/utils"
)

func main() {
	cli := &cli.CliApp{
		Config: &utils.Config{},
		Args:   os.Args,
	}
	cli.Initialize()
	cli.Run()
}
