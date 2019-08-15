package auth

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var groupSQLQuery = `
	CREATE TABLE IF NOT EXISTS "group" (
		"id" serial primary key,
		"name" varchar(150) not null
	);`

// Group is the group entity schema
type Group struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

// GetUsers will show all group users
func (g *Group) GetUsers(dbCon *sqlx.DB) []User {
	// TODO: make query and select data from group_users
	usersData := []User{}
	err := dbCon.Select(&usersData, `SELECT "id","email","username","password","last_login","first_name","last_name","is_active" FROM "group_users" ORDER BY "id" ASC;`)
	if err != nil {
		log.Fatalf("Error while unmarshal to struct users data: %v", err)
	}
	return usersData
}

// NewGroup is the Group type factory function
func NewGroup(name string) *Group {
	return &Group{
		Name: name,
	}
}
