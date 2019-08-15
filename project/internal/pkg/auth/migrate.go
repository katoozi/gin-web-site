package auth

import "github.com/jmoiron/sqlx"

// MigrateTables will execute all sql queries
func MigrateTables(db *sqlx.DB) {
	query := userSQLQuery + sessionSQLQuery
	db.MustExec(query)
}
