package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// User will have the user table schema
type User struct {
	ID        int       `db:"id" sqltools:"id"`
	FirstName string    `db:"first_name" sqltools:"first_name"`
	LastName  string    `db:"last_name" sqltools:"last_name"`
	Password  string    `db:"password" sqltools:"password"`
	LastLogin time.Time `db:"last_login" sqltools:"last_login"`
	Username  string    `db:"username" sqltools:"username"`
	IsActive  bool      `db:"is_active" sqltools:"is_active" default:"1"`
	Email     string    `db:"email" sqltools:"email"`
}

// NewUser is the User type factory function
func NewUser(firstName, lastName, username, email, password string, lastLogin time.Time) *User {
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
		IsActive:  true,
	}
}

// NewUserIsActive is the User type factory function
func NewUserIsActive(firstName, lastName, username, email, password string, lastLogin time.Time, isActive bool) *User {
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
		IsActive:  isActive,
	}
}

// GenerateInsertQuery will generate sql insert query for postgresql
func (u *User) GenerateInsertQuery() string {
	schema := `INSERT INTO "user" (first_name,last_name,email,username,password,last_login,is_active) VALUES ('%s','%s','%s','%s','%s','%s','%t');`
	return fmt.Sprintf(
		schema,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Username,
		u.Password,
		u.LastLogin.Format(time.UnixDate),
		u.IsActive,
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
		"id" serial not null PRIMARY KEY,
		"first_name" text,
		"last_name" text,
		"password" text not null,
		"last_login" timestamptz,
		"username" text unique not null,
		"email" text,
		"is_active" boolean not null default '1'
	);
	`
	db.MustExec(userTableSQLQuery)
}
