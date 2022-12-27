package config

import (
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

var DebugMode bool

func init() {
	DebugMode = os.Getenv("DEBUG") == "true"
}

type Config struct {
	ClientID     []string `yaml:"client_id"`
	ClientSecret []string `yaml:"client_secret"`
	AuthURL      []string `yaml:"auth_url"`
	TokenURL     []string `yaml:"token_url"`
	AppPort      int      `yaml:"app_port"`
	OAuth2Config *oauth2.Config
}

func New(filePath string) *Config {
	if filePath == "" {
		filePath = "env.yml"
	}
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

	var config Config

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Println("Unable to parse contents of YAML config file:", filePath)
		log.Fatalln(err)
	}
	if DebugMode {
		log.Println("Loaded config:")
		config.Print()
	}

	validateConfig(&config)

	return &config
}

func validateConfig(c *Config) {
	var fieldsMissing []string
	if c.ClientID == "" {
		fieldsMissing = append(fieldsMissing, "client_id")
	}
	if c.ClientSecret == "" {
		fieldsMissing = append(fieldsMissing, "client_secret")
	}
	if len(fieldsMissing) > 0 {
		log.Println("The follow fields appear to be missing from your config file:")
		for _, configKey := range fieldsMissing {
			log.Println("-", configKey)
		}
		log.Fatalln("Please ensure all required config values are present.")
	}
}

// Print outputs the loaded client struct.
func (c *Config) Print() {
	log.Println("Client ID:    ", c.ClientID)
	log.Println("Client Secret:", c.ClientSecret)
}
