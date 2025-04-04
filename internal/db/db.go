package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	var err error

	db, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		return nil, err
	}

	// Crear tablas va aquí
	_, err = db.Exec(` SELECT 1 `)

	if err != nil {
		return nil, fmt.Errorf("Error creating tables: %s\n", err)
	}

	return db, nil
}
