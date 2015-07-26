package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20150114025706(txn *sql.Tx) {
	txn.Exec(`
		CREATE TABLE users (
			id SERIAL NOT NULL PRIMARY KEY,
			first_name varchar(255) NOT NULL,
			last_name varchar(255) NOT NULL,
			company varchar(255) NOT NULL,
			email varchar(255) UNIQUE NOT NULL,
			password_digest varchar(60) NOT NULL,
			role_id integer NOT NULL,
			last_seen timestamp DEFAULT current_timestamp,
			updated_at timestamp DEFAULT current_timestamp,
			created_at timestamp DEFAULT current_timestamp
		)
	`)
}

// Down is executed when this migration is rolled back
func Down_20150114025706(txn *sql.Tx) {
	txn.Exec(`
		DROP TABLE users
	`)
}
