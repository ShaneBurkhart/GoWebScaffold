package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20150605192632(txn *sql.Tx) {
	txn.Exec(`
		CREATE TABLE roles (
			id SERIAL NOT NULL PRIMARY KEY,
			name varchar(255) NOT NULL
		)
	`)
	txn.Exec(`
		INSERT INTO roles (
			name
		) VALUES
			('admin'),
			('manager'),
			('viewer')
	`)
}

// Down is executed when this migration is rolled back
func Down_20150605192632(txn *sql.Tx) {
	txn.Exec(`
		DROP TABLE roles
	`)
}
