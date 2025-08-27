package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	var connectString = "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		panic(err.Error())
	}
	return db
}
