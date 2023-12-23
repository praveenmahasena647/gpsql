package app

import (
	dbs "github.com/praveenmahasena647/gpsql/cmd/app/DBs"
	"github.com/praveenmahasena647/gpsql/cmd/app/configs"
)

func Exc() error {
	var YAMLstruct, yamlErr = configs.YAMLfmt()
	if yamlErr != nil {
		return yamlErr
	}

	return dbs.Exc(YAMLstruct)
}
