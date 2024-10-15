package users

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// Create a table using an existing database connection.
func CreateTable(db *sql.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        age INTEGER
    );`
	_, err := db.Exec(createTableSQL)
	return err
}
