package dbs

import (
	"errors"
	"os"
	"uri/cmd/internal/models"
)

func Runfiles(p string, yamlConfig models.YamlStruct) error {
	if len(os.Args) == 1 {
		return errors.New("did not provide the correct args")
	}
	switch os.Args[1] {
	case "up":
		return dbUp(p + "/" + yamlConfig.Upfile)
	case "down":
		return dbDown(p + "/" + yamlConfig.Downfile)
	}
	return errors.New("did not give the commend up or not")
}
