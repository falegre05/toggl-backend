package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// Database connection
var _db *sqlx.DB

func init() {
	// Open SQLite database
	db := GetDBConnection()

	// Create tables if they don't exist
	db.MustExec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE,
			password_hash TEXT
		);

		CREATE TABLE IF NOT EXISTS questions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			body TEXT,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id),
			UNIQUE(body, user_id)
		);

		CREATE TABLE IF NOT EXISTS options (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			body TEXT,
			correct BOOLEAN,
			question_id INTEGER,
			FOREIGN KEY (question_id) REFERENCES questions(id),
			UNIQUE(body, question_id)
		);
	`)
}

func GetDBConnection() *sqlx.DB {
	if _db == nil {
		var err error
		_db, err = sqlx.Connect("sqlite3", "toggl-backend.db")
		if err != nil {
			log.Fatal(err)
		}
	}

	return _db
}
