package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./series.db")

	if err != nil {
		log.Fatal(err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS AnimeManga(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		type TEXT,
		total INTEGER,
		progress INTEGER DEFAULT 0,
		image TEXT
	);`

	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal(err)
	}
}
