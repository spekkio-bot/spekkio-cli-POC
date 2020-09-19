package cli

import (
	"fmt"

	"github.com/spekkio-bot/spekkio-cli/src/config"
)

type CliApp struct {
	Config *config.Config
	Args   []string
}

func (c *CliApp) Initialize() {
	c.Config.Initialize()
}

func (c *CliApp) Run() {
	if c.SelectedPing() {
		c.Ping()
	} else if c.SelectedConfigure() {
		c.Configure()
	} else if c.SelectedDefault() {
		fmt.Println("i am the master of war!")
	} else {
		fmt.Println("err: invalid args!")
	}
}
