package config

import (
	"encoding/json"
	"os"
	"sync"
)

// Enum values
const (
	EnvironmentDevelopment string = "dev"
	EnvironmentProduction  string = "production"

	ServiceEnabled  string = "Yes"
	ServiceDisabled string = "No"
)

// Configuration works with configuration of this project
type Configuration struct {
	Environment string
	Port        string
	Database    struct {
		Path string
		Port string
		Name string
	}
	SMTP struct {
		EnableService string
		Host          string
		Port          string
		UserName      string
		Password      string
	}
	RandomSeedOffset int64
}

var instance *Configuration
var once sync.Once

// GetConfiguration will return the configuration with go object
func GetConfiguration(configFilePath string) *Configuration {
	once.Do(func() {
		file, _ := os.Open(configFilePath)
		defer file.Close()
		decoder := json.NewDecoder(file)
		instance = &Configuration{}
		err := decoder.Decode(&instance)
		if err != nil {
			panic("Configuration parsing error!!!")
		}
	})

	return instance
}
