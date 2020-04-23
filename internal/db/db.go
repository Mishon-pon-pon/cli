package db

var config *Config

// DB ...
type DB struct {
	Connection struct {
		ConnectionString string `json:"connectionString"`
	}
}

// Config ...
type Config struct {
	DataBaseURL string `json:"connectionString"`
	DBManager   string `json:"db_manager"`
}

// NewConfig ...
func NewConfig() *Config {
	if config != nil {
		return config
	}
	config = &Config{}
	return config
}
