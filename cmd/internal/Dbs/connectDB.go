package dbs

import (
	"database/sql"
	"fmt"
	"uri/cmd/internal/models"

	_ "github.com/lib/pq"
)

var (
	C *sql.DB
)

func DBconnection(m models.YamlStruct) error {
	var connectionErr error
	var connectionUri = fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v", m.User, m.Password, m.Dbname, m.Host, m.Port, m.Sslmode)
	C, connectionErr = sql.Open("postgres", connectionUri)
	if connectionErr != nil {
		return connectionErr
	}
	return C.Ping()
}
