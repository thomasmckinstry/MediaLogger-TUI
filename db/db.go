/*
Copyright © 2025 Thomas McKinstry thomas.g.mckinstry@protonmail.com
*/

package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetDB() *sql.DB {
	var err error

	if db != nil {
		return db
	}

	db, err := sql.Open("sqlite3", "./media.db")
	if err != nil {
		log.Fatal("Unable to open database:", err)

		os.Exit(1)
	}
	init_db(db)

	return db
}
