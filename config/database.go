package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := ""
	dbName := "managementtoko"

	db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"/@"+dbName)
	return
}
