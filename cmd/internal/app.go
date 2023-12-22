package internal

import (
	"fmt"
	dbs "uri/cmd/internal/Dbs"
	handlefile "uri/cmd/internal/handleFile"
)

func Start() error {
	var wdDir, dirErr = handlefile.GetWD()
	if dirErr != nil {
		return dirErr
	}
	var fileBytes, fileErr = handlefile.ReadYaml(wdDir)
	if fileErr != nil {
		fmt.Println("config File Err")
		return fileErr
	}
	var yamlStruct, yamlErr = handlefile.FmtYaml(fileBytes)
	if yamlErr != nil {
		fmt.Println("error during Yaml file Parsing")
		return yamlErr
	}

	var dbConnectionErr error = dbs.DBconnection(yamlStruct)

	if dbConnectionErr != nil {
		return dbConnectionErr
	}
	var eventErr error = dbs.Runfiles(wdDir, yamlStruct)
	if eventErr != nil {
		return eventErr
	}

	return nil
}
