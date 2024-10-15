package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yeungon/tuhuebot/internal/database/sqlite/users"
	_ "modernc.org/sqlite"
)

// Initialize and open the database connection.
func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "file:tuhuebot.db")
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to sqlite")

	// Create the users table if it doesn't exist.
	if err := users.CreateTable(db); err != nil {
		log.Fatal("Failed to create table:", err)
	}
	////
	return db, nil
}
