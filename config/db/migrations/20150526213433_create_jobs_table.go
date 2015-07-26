package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20150526213433(txn *sql.Tx) {
	txn.Exec(`
		CREATE TABLE jobs (
			id SERIAL NOT NULL PRIMARY KEY,
			name varchar(255) NOT NULL,
			user_id integer NOT NULL
		)
	`)
}

// Down is executed when this migration is rolled back
func Down_20150526213433(txn *sql.Tx) {
	txn.Exec(`
		DROP TABLE jobs
	`)
}
