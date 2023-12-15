package db

import (
	"database/sql"
	_ "database/sql"
	_ "pract7/types"

	_ "github.com/go-sql-driver/mysql"
)

func connect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "test"
	dbPass := "test_password"
	dbName := "api_pract_7_database"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
