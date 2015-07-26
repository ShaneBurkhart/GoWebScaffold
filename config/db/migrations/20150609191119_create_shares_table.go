package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20150609191119(txn *sql.Tx) {
	txn.Exec(`
		CREATE TABLE shares (
			id SERIAL NOT NULL PRIMARY KEY,
			job_id integer NOT NULL,
			sharer_id integer NOT NULL,
			user_id integer NOT NULL
		)
	`)
}

// Down is executed when this migration is rolled back
func Down_20150609191119(txn *sql.Tx) {
	txn.Exec(`
		DROP TABLE shares
	`)
}
