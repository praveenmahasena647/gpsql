package dbs

import (
	"errors"
	"os"

	"github.com/praveenmahasena647/gpsql/cmd/app/models"
)

// this function calls and passes the up || down file path
func Exc(y models.YAMLstruct) error {
	var connectionErr = Connection(y)
	if connectionErr != nil {
		return connectionErr
	}
	if len(os.Args) < 1 {
		return errors.New("did not provide proper cmd args")
	}
	var wd, _ = os.Getwd()
	switch os.Args[1] {
	case "up":
		return helper(wd + "/" + y.Updir)
	case "down":
		return helper(wd + "/" + y.Downdir)
	default:
		return errors.New("did not provide proper cmd args")
	}
}
