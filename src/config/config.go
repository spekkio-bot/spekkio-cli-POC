package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	DEFAULT_SERVER_URL = "https://5ila6fw37k.execute-api.us-west-1.amazonaws.com/api"
)

type Config struct {
	ServerUrl string
}

func (c *Config) Initialize() {
	configfile, err := ioutil.ReadFile(spekkiofilePath())
	if err != nil {
		fmt.Printf("no spekkiofile file found, creating spekkiofile and using defaults\n")
		initSpekkiofile()
		c.UseDefault()
	} else {
		configmap := parseSpekkiofile(configfile)
		c.ServerUrl = configmap["SERVER_URL"]
	}
}

func (c *Config) UseDefault() {
	c.ServerUrl = DEFAULT_SERVER_URL
}

func parseSpekkiofile(configfileBytes []byte) map[string]string {
	configfile := string(configfileBytes)
	configmap := make(map[string]string)

	items := strings.Split(configfile, "\n")
	for _, item := range items {
		if item == "" {
			continue
		}
		itemSlice := strings.Split(item, "=")
		configmap[itemSlice[0]] = itemSlice[1]
	}

	return configmap
}

func spekkiofilePath() string {
	configpath, _ := os.UserHomeDir()
	configpath += "/.spekkio/spekkiofile"
	return configpath
}

func spekkiofileDirPath() string {
	configdirpath, _ := os.UserHomeDir()
	configdirpath += "/.spekkio"
	return configdirpath
}

func initSpekkiofile() {
	err := os.Mkdir(spekkiofileDirPath(), 0755)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Create(spekkiofilePath())
	if err != nil {
		log.Fatal(err)
	}
	var contents []byte
	contents = append(contents, []byte("SERVER_URL=")...)
	contents = append(contents, []byte(DEFAULT_SERVER_URL)...)
	contents = append(contents, []byte("\n")...)
	err = ioutil.WriteFile(spekkiofilePath(), contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
