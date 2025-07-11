package db

import (
	"database/sql"
)

var DB *sql.DB

func InitService(database *sql.DB) {
	DB = database
}
