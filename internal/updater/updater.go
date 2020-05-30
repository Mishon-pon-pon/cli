package updater

import (
	"fso/internal/fsomodule"
	"fso/internal/fsoservice"
)

// Config ...
type Config struct {
	PathFrom string `json:"pathFrom"`
	PathIn   string `jsoon:"pathIn"`
}

// IUpdate - interface for cocrete implementation update
type IUpdate interface {
	Update(string, string, string) error
}

// NewUpdater - constructor for Updater type factoryMethod
func NewUpdater(updaterTargetName string) IUpdate {
	var c IUpdate
	switch updaterTargetName {
	case "Service":
		c = fsoservice.NewService()
	case "Module":
		c = fsomodule.NewModule()
	}

	return c
}
