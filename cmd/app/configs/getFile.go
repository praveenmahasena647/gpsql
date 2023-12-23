package configs

import (
	"os"

	"github.com/praveenmahasena647/gpsql/cmd/app/models"
	"gopkg.in/yaml.v3"
)

func YAMLfmt() (models.YAMLstruct, error) {
	var yamlStruct models.YAMLstruct
	var workDir, workDirErr = os.Getwd()
	if workDirErr != nil {
		return yamlStruct, workDirErr
	}
	var fileBytes, fileErr = os.ReadFile(workDir + "/gpsql.yaml")
	if fileErr != nil {
		return yamlStruct, fileErr
	}
	yamlStruct = models.YAMLstruct{}
	var err = yaml.Unmarshal(fileBytes, &yamlStruct)
	return yamlStruct, err
}
