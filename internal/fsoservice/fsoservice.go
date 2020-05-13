package fsoservice

// Service ...
type Service struct {
	/* folderThatContainConfig нужен для того чтобы
	 * при удалении папок знать в какийх папках
	 * хроняться конфиги .config
	 * используется в DeleteModuel и в CopyModule
	 */
	folderThatContainConfig map[string][]string
}

// NewService - constructor Module
func NewService() *Service {
	return &Service{folderThatContainConfig: map[string][]string{}}
}

// CopyError ...
type CopyError struct {
	err error
	msg string
}

func (c *CopyError) Error() string { return c.msg }
