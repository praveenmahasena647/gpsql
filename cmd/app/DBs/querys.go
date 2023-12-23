package dbs

import "os"

// this function reads all the files inside of provided up || down path
// The files are sorted by name this might give somekind of err while doing the Query
func helper(path string) error {
	var files, fileErr = os.ReadDir(path)
	if fileErr != nil {
		return fileErr
	}
	for _, v := range files {
		var fileBytes, byteErr = os.ReadFile(path + "/" + v.Name())
		if byteErr != nil {
			return byteErr
		}
		if _, err := c.Query(string(fileBytes)); err != nil {
			return err
		}
	}
	return nil
}
