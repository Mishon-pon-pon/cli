package modules

// Config ...
type Config struct {
	PathFrom string `json:"pathFrom"`
	PathIn   string `jsoon:"pathIn"`
}

// Module ...
type Module struct {
	ConfigFiles map[string][]string
}

// NewModule - constructor Module
func NewModule() *Module {
	return &Module{}
}

// CopyError ...
type CopyError struct {
	err error
	msg string
}

func (c *CopyError) Error() string { return c.msg }
