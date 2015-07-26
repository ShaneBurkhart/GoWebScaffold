package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20150528031209(txn *sql.Tx) {
	txn.Exec(`
		CREATE TABLE plans (
			id SERIAL NOT NULL PRIMARY KEY,
			name varchar(255) NOT NULL,
			num integer NOT NULL,
			job_id integer NOT NULL
		)
	`)
}

// Down is executed when this migration is rolled back
func Down_20150528031209(txn *sql.Tx) {
	txn.Exec(`
		DROP TABLE plans
	`)
}
