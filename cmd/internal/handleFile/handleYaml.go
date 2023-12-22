package handlefile

import (
	"os"
	"uri/cmd/internal/models"

	"gopkg.in/yaml.v3"
)

func GetWD() (string, error) {
	return os.Getwd()
}

func ReadYaml(wordDir string) ([]byte, error) {
	return os.ReadFile(wordDir + "/uri.yaml")
}

func FmtYaml(b []byte) (models.YamlStruct, error) {
	var vals = models.YamlStruct{}
	var err = yaml.Unmarshal(b, &vals)
	return vals, err
}
