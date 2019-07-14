package models

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// User will have the user table schema
type User struct {
	ID        int    `db:"id" sqltools:"id"`
	FirstName string `db:"first_name" sqltools:"first_name"`
	LastName  string `db:"last_name" sqltools:"last_name"`
	Password  string `db:"password" sqltools:"password"`
	LastLogin string `db:"last_login" sqltools:"last_login"`
	Username  string `db:"username" sqltools:"username"`
	Email     string `db:"email" sqltools:"email"`
}

// NewUser is the User type factory function
func NewUser(firstName, lastName, username, email, password, lastLogin string) *User {
	hashpassword, err := generate(password)
	if err != nil {
		log.Fatalf("Error while encrypting password: %v", err)
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  hashpassword,
		LastLogin: lastLogin,
	}
}

// GenerateInsertQuery will generate sql insert query for postgresql
func (u *User) GenerateInsertQuery() string {
	schema := `INSERT INTO "user" (first_name,last_name,email,username,password,last_login) VALUES ('%s','%s','%s','%s','%s','%s');`
	return fmt.Sprintf(
		schema,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Username,
		u.Password,
		u.LastLogin,
	)
}

func generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

//Compare string to generated hash
func (u *User) Compare(s string) error {
	incoming := []byte(s)
	return bcrypt.CompareHashAndPassword([]byte(u.Password), incoming)
}

// GetUser will fetch user from db by username
func GetUser(username string, dbCon *sqlx.DB) *User {
	query := fmt.Sprintf(`SELECT * from "user" WHERE username='%s';`, username)
	userObj := User{}
	err := dbCon.Get(&userObj, query)
	if err != nil {
		log.Printf("Error while get user from db:%v", err)
		return nil
	}
	return &userObj
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
}
