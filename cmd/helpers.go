package cmd

import "os"

func generateYaml() error {
	var wd, wdErr = os.Getwd()
	if wdErr != nil {
		return wdErr
	}
	var configStr = `
		user: # dbUserName
		password: #enter password
		dbname: # db name you wanna connect with
		host: #db Host
		port: #db port
		sslmode: #sslmode
		updir: #relative path to the up files in your current folder example here ./cmd/up-files
		downdir: #relative path to the down folder from your current working dir  example here ./cmd/down-files
	`
	return os.WriteFile(wd+"/"+"gpsql.yaml", []byte(configStr), os.ModePerm)
}
