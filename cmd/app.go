package cmd

import (
	"errors"
	"os"

	"github.com/praveenmahasena647/gpsql/cmd/db"
)

func Start() error {
	if len(os.Args) == 1 {
		return errors.New("did not provite enought args")
	}

	if os.Args[1] == "generate" {
		return generateYaml()
	}

	var configs = db.New()
	var initErr = configs.Init()
	if initErr != nil {
		return initErr
	}

	switch os.Args[1] {
	case "up":
		return configs.Up()
	case "down":
		return configs.Down()
	default:
		return errors.New("did not provide the correct args")
	}
}
