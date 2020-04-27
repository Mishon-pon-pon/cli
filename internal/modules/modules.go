package modules

// Config ...
type Config struct {
	PathFrom string `json:"pathFrom"`
	PathIn   string `jsoon:"pathIn"`
}

// Module ...
type Module struct {
	/* folderThatContainConfig нужен для того чтобы
	 * при удалении папок знать в какийх папках
	 * хроняться конфиги .config
	 * используется в DeleteModuel и в CopyModule
	 */
	folderThatContainConfig map[string][]string
}

// NewModule - constructor Module
func NewModule() *Module {
	return &Module{folderThatContainConfig: map[string][]string{}}
}

// CopyError ...
type CopyError struct {
	err error
	msg string
}

func (c *CopyError) Error() string { return c.msg }
