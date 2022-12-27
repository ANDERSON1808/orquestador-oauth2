package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"orquestador/basic/models"
	"os"

	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

var DebugMode bool

func init() {
	DebugMode = os.Getenv("DEBUG") == "true"
}

type Config struct {
	AppPort      int `yaml:"app_port"`
	Services     []*models.Services
	OAuth2Config *oauth2.Config
}

func New(filePath string) *Config {
	var config Config
	if filePath == "" {
		fmt.Println("start without configuration file")
	} else {
		// Check that the file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Println("Config file does not exist:", filePath)
			log.Fatalln(err)
		}
		configFile, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Println("An error occurred while trying to read the config file:", filePath)
			log.Fatalln(err)
		}

		err = yaml.Unmarshal(configFile, &config)
		if err != nil {
			log.Println("Unable to parse contents of YAML config file:", filePath)
			log.Fatalln(err)
		}
		return &config
	}
	return &config
}
