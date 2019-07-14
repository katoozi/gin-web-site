package models

import "github.com/jmoiron/sqlx"

// MigrateTables create tables if they not exists.
func MigrateTables(db *sqlx.DB) {
	// create user table
	userTableSQLQuery := `
	CREATE TABLE IF NOT EXISTS "user" (
		"id" serial not null PRIMARY KEY,
		"first_name" varchar(30),
		"last_name" varchar(150),
		"password" varchar(130) not null,
		"last_login" timestamptz,
		"date_joined" timestamptz default now(),
		"username" varchar(150) unique not null,
		"email" varchar(254),
		"is_active" boolean not null default 'true',
		"is_staff" boolean not null default 'false',
		"is_superuser" boolean not null default 'false'
	);
	CREATE TABLE IF NOT EXISTS "session" (
		"session_key" varchar(40) not null primary key,
		"session_data" text not null,
		"expire_date" timestamptz not null
	);
	`
	db.MustExec(userTableSQLQuery)
}
