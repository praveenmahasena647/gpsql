package db

import (
	"database/sql"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DBModel struct {
	Wd       string
	User     string `yaml:user`
	Password string `yaml:password`
	Dbname   string `yaml:dbname`
	Host     string `yaml:host`
	Port     string `yaml:port`
	Sslmode  string `yaml:sslmode`
	Updir    string `yaml:updir`
	Downdir  string `yaml:downdir`
}

func New() *DBModel {
	return &DBModel{}
}

func (d *DBModel) Init() error {
	var wd, wdErr = os.Getwd()
	if wdErr != nil {
		return wdErr
	}
	d.Wd = wd
	var fbytes, fErr = os.ReadFile(wd + "/" + "gpsql.yaml")
	if fErr != nil {
		return fErr
	}
	return yaml.Unmarshal(fbytes, d)
}

func (d *DBModel) Up() error {
	var connection, connectionErr = connect(d)
	if connectionErr != nil {
		return connectionErr
	}
	return query(d.Wd+"/"+d.Updir, connection)
}

func (d *DBModel) Down() error {
	var connection, connectionErr = connect(d)
	if connectionErr != nil {
		return connectionErr
	}
	return query(d.Wd+"/"+d.Downdir, connection)
}

func query(path string, connection *sql.DB) error {
	var someThing, err = os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, v := range someThing {
		var bytes, byteErr = os.ReadFile(path + "/" + v.Name())
		if byteErr != nil {
			return byteErr
		}
		if _, err := connection.Query(string(bytes)); err != nil {
			return err
		}
	}
	return nil
}

func connect(d *DBModel) (*sql.DB, error) {
	var dns = fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v", d.User, d.Password, d.Dbname, d.Host, d.Port, d.Sslmode)
	var db, dbErr = sql.Open("postgres", dns)
	if dbErr != nil {
		return nil, dbErr
	}
	return db, db.Ping()
}
