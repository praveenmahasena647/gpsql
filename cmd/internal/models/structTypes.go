package models

type YamlStruct struct {
	User     string `yaml:user`
	Password string `yaml:password`
	Dbname   string `yaml:dbname`
	Host     string `yaml:host`
	Port     string `yaml:post`
	Sslmode  string `yaml:sslmode`
	Upfile   string `yaml:upfile`
	Downfile string `yaml:downfile`
}
