package models

type YAMLstruct struct {
	User     string `yaml:user`
	Password string `yaml:password`
	DBname   string `yaml:dbname`
	Host     string `yaml:host`
	Port     string `yaml:port`
	Sslmode  string `yaml:sslmode`
	Updir    string `yaml:updir`
	Downdir  string `yaml:downdir`
}
