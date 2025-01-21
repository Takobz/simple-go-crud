package configuration

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
