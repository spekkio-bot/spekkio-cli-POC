package utils

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

// Config defines an app's configuration
type Config struct {
	ServerUrl string
}

// Initialize will read from spekkiofile and initialize the Config object's values
// If no spekkiofile is present, it will create one
func (c *Config) Initialize() {
	configfile, err := ioutil.ReadFile(SpekkiofilePath())
	if err != nil {
		fmt.Printf("no spekkiofile file found, creating spekkiofile and using defaults\n")
		initSpekkiofile()
		c.UseDefault()
	} else {
		configmap := ParseSpekkiofile(configfile)
		c.ServerUrl = configmap["SERVER_URL"]
	}
}

// UseDefault will apply the default constant values to the Config object
func (c *Config) UseDefault() {
	c.ServerUrl = DEFAULT_SERVER_URL
}

// ParseSpekkiofile parses spekkiofile into a map of string-string key-value pairs
func ParseSpekkiofile(configfileBytes []byte) map[string]string {
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

// SpekkiofilePath returns the directory path of spekkiofile
func SpekkiofilePath() string {
	configpath, _ := os.UserHomeDir()
	configpath += "/.spekkio/spekkiofile"
	return configpath
}

// SpekkiofileDirPath returns the directory path of the folder containing spekkiofile
func SpekkiofileDirPath() string {
	configdirpath, _ := os.UserHomeDir()
	configdirpath += "/.spekkio"
	return configdirpath
}

func initSpekkiofile() {
	err := os.Mkdir(SpekkiofileDirPath(), 0755)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Create(SpekkiofilePath())
	if err != nil {
		log.Fatal(err)
	}
	var contents []byte
	contents = append(contents, []byte("SERVER_URL=")...)
	contents = append(contents, []byte(DEFAULT_SERVER_URL)...)
	contents = append(contents, []byte("\n")...)
	err = ioutil.WriteFile(SpekkiofilePath(), contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
