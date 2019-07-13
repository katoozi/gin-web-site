package website

import (
	"github.com/jmoiron/sqlx"
)

var dbCon *sqlx.DB

// User will have the user table schema
type User struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Password  string `db:"password"`
	LastLogin string `db:"last_login"`
	Username  string `db:"username"`
	Email     string `db:"email"`
}

// MigrateTables create tables if they not exists.
func MigrateTables(db *sqlx.DB) {
	// create user table
	userTableSQLQuery := `
	CREATE TABLE IF NOT EXISTS "user" (
		"id" serial,
		"first_name" text,
		"last_name" text,
		"password" text,
		"last_login" timestamp,
		"username" text,
		"email" text,
		PRIMARY KEY("id")
	);
	`

	db.MustExec(userTableSQLQuery)

	dbCon = db
}
