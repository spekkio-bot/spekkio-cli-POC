package cli

import (
	"testing"

	"github.com/spekkio-bot/spekkio-cli/src/config"
)

func TestNoArgs(t *testing.T) {
	c := CliApp{
		Config: &config.Config{
			ServerUrl: "shouldWork",
		},
		Args: []string{"test"},
	}

	c.Run()
}

func TestInvalidArgs(t *testing.T) {
	c := CliApp{
		Config: &config.Config{
			ServerUrl: "shouldAlsoWork",
		},
		Args: []string{"this", "is", "invalid"},
	}

	c.Run()
}
