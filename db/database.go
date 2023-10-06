package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfiguraton struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

func Connect(cnf DatabaseConfiguraton) (*sql.DB, error) {
	//username:password@tcp(host:port)/database
	urlConnection := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.DatabaseName,
	)
	db, err := sql.Open("mysql", urlConnection)
	return db, err
}
