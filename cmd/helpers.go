package cmd

import "os"

func generateYaml() error {
	var wd, wdErr = os.Getwd()
	if wdErr != nil {
		return wdErr
	}
	var configStr = "user: # dbUserName\npassword: #enter password\ndbname: # db name you wanna connect with\nhost: #db Host\nport: #db port\nsslmode: #sslmode\nupdir: #relative path to the up files in your current folder example here ./cmd/up-files\ndowndir: #relative path to the down folder from your current working dir  example here ./cmd/down-files"
	return os.WriteFile(wd+"/"+"gpsql.yaml", []byte(configStr), os.ModePerm)
}
