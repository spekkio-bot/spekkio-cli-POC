package cli

import (
	"fmt"

	"github.com/spekkio-bot/spekkio-cli/src/config"
)

type CliApp struct {
	Config *config.Config
}

func (c *CliApp) Initialize() {
	c.Config.Initialize()
}
