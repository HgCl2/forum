package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func Init() {
	var err error
	database, err := sql.Open("sqlite3", "file:forum.db")
	if err != nil {
		panic("failed to connect database")
	}

	statement, err := database.Prepare("PRAGMA foreign_keys = 1")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	// CreateTables()
	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS users " +
		"(id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, nickname TEXT," +
		"email TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS posts " +
		"(id INTEGER PRIMARY KEY, message TEXT, author TEXT, email TEXT, " +
		"like INTEGER DEFAULT 0, dislike INTEGER DEFAULT 0, user_id INTEGER," +
		"FOREIGN KEY (user_id) REFERENCES users(id))")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS categories " +
		"(id INTEGER PRIMARY KEY, name TEXT, post_id INTEGER," +
		"FOREIGN KEY (post_id) REFERENCES posts(id))")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS comments " +
		"(id INTEGER PRIMARY KEY, content TEXT, author TEXT, post_id INTEGER," +
		"FOREIGN KEY (post_id) REFERENCES posts(id))")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
}

// func CreateTables() {

// }
