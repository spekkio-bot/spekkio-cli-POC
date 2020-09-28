package cli

import (
	"testing"

	"github.com/spekkio-bot/spekkio-cli/src/utils"
)

func TestNoArgs(t *testing.T) {
	c := CliApp{
		Config: &utils.Config{
			ServerUrl: "shouldWork",
		},
		Args: []string{"test"},
	}

	c.Run()
}

func TestInvalidArgs(t *testing.T) {
	c := CliApp{
		Config: &utils.Config{
			ServerUrl: "shouldAlsoWork",
		},
		Args: []string{"this", "is", "invalid"},
	}

	c.Run()
}
