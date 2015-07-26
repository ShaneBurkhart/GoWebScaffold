package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func SetupDB() error {
	var err error
	DB, err = sql.Open("postgres", "host=pg user=postgres password=postgres dbname=mydb sslmode=disable")
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		DB.Close()
		return err
	}
	return nil
}

func VerifyDB() error {
	return DB.Ping()
}

func CloseDB() {
	DB.Close()
}
