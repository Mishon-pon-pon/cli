package repo

var conf *Config

// Config ...
type Config struct {
	RemotePath string `json:"remotePath"`
	DevPath    string `json:"devPath"`
	TestPath   string `json:"testPath"`
}

// NewConfig ...
func NewConfig() *Config {
	if conf != nil {
		return conf
	}
	conf = &Config{}
	return conf
}
