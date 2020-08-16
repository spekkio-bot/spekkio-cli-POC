package cli

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func (c *CliApp) Ping() {
	res, err := http.Get(c.Config.ServerUrl)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		fmt.Println("spekkio is online.")
	} else {
		fmt.Printf("err: spekkio responded with status %d.\n", res.StatusCode)
		os.Exit(1)
	}
}
