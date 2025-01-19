package configuration

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type ConfigurationOptions struct {
	Postgres DatabaseOptions `json:"postgres"`
}

type DatabaseOptions struct {
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func GetConfiguration() ConfigurationOptions {
	configurationOptions := ConfigurationOptions{}
	configuration := ReadConfigurationFileFromHost()
	err := json.Unmarshal(configuration, &configurationOptions)
	if err != nil {
		panic(err)
	}

	return configurationOptions
}

func ReadConfigurationFileFromHost() []byte {
	//path relative to caller i.e main.go
	absPath, err := filepath.Abs("./appsettings.json")
	if err != nil {
		panic(err)
	}

	configurationFile, err := os.Open(absPath)
	if err != nil {
		panic(err)
	}
	defer configurationFile.Close()

	configuration, err := io.ReadAll(configurationFile)
	if err != nil {
		panic(err)
	}
	return configuration
}
