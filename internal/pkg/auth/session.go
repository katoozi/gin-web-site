package auth

import "time"

var sessionSQLQuery = `
	CREATE TABLE IF NOT EXISTS "session" (
		"session_key" varchar(40) not null primary key,
		"session_data" text not null,
		"expire_date" timestamptz not null
	);`

// Session is the session entity schema
type Session struct {
	SessionKey  string    `db:"session_key"`
	SessionData string    `db:"session_data"`
	ExpireDate  time.Time `db:"expire_date"`
}

// NewSession is the Session type factory function
func NewSession(sessionData string, expireTime time.Time) *Session {
	return &Session{
		ExpireDate: expireTime,
	}
	// TODO: encrypt sessionData and generate session key for the session later
}
