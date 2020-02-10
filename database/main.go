package database

import (
	"database/sql"
)

// StartDatabase export an instance of MySql
func StartDatabase(driverName string, dataSourceName string) (db *sql.DB, err error) {

	db, err = sql.Open(driverName, dataSourceName)

	return db, err
}
