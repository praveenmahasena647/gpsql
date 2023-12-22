package dbs

import (
	"os"
)

func dbUp(p string) error {
	if _, err := os.Stat(p); err != nil {
		return err
	}
	var files, fileErr = os.ReadDir(p)
	if fileErr != nil {
		return fileErr
	}
	for _, f := range files {
		var fileBytes, byteErr = os.ReadFile(p + "/" + f.Name())
		if byteErr != nil {
			return fileErr
		}
		if _, err := C.Exec(string(fileBytes)); err != nil {
			return err
		}
	}
	return nil
}
