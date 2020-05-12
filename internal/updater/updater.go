package updater

import (
	"fso/internal/fsomodule"
	"fso/internal/fsoservice"
)

type Config struct {
	PathFrom string `json:"pathFrom"`
	PathIn   string `jsoon:"pathIn"`
}

type ICopy interface {
	Copy(string, string, string) error
}

func NewUpdater(updaterTargetName string) ICopy {
	var c ICopy
	switch updaterTargetName {
	case "Service":
		c = fsoservice.NewService()
	case "Module":
		c = fsomodule.NewModule()
	}

	return c
}
