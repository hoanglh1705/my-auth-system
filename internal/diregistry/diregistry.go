package diregistry

import (
	"my-auth-service/config"
	"my-auth-service/internal/helper/dihelper"

	"github.com/sarulabs/di"
)

const (
	ConfigDIName string = "Config"
)

func BuildDIContainer() {
	initBuilder()
	dihelper.BuildLibDIContainer()
}

func GetDependency(name string) interface{} {
	return dihelper.GetLibDependency(name)
}

func initBuilder() {
	dihelper.ConfigsBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  ConfigDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg, err := config.Load()
				return cfg, err
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})

		return arr
	}
}
