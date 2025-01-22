package configuration

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

func GetConfiguration() ConfigurationOptions {
	configurationOptions := ConfigurationOptions{}
	configuration := readConfigurationFileFromHost()
	err := json.Unmarshal(configuration, &configurationOptions)
	if err != nil {
		panic(err)
	}

	return configurationOptions
}

func GetPostgresConnectionString() string {
	configOptions := GetConfiguration()

	//produces a string like "postgres://user:password@host/database?sslmode=disable"
	return "postgres://" +
		configOptions.Postgres.User + ":" +
		configOptions.Postgres.Password + "@" +
		configOptions.Postgres.Host + "/" +
		configOptions.Postgres.Database + "?sslmode=disable"
}

/*
write about how method can be made private by making them lowercase.
*/
func readConfigurationFileFromHost() []byte {
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
