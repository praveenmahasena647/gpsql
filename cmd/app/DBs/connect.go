package dbs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/praveenmahasena647/gpsql/cmd/app/models"
)

var (
	c *sql.DB
)

func Connection(y models.YAMLstruct) error {
	var err error
	var connectionUri = fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v", y.User, y.Password, y.DBname, y.Host, y.Port, y.Sslmode)
	c, err = sql.Open("postgres", connectionUri)
	if err != nil {
		return err
	}
	return c.Ping()
}
