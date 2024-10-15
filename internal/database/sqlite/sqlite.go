package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func SQLite() {
	// Open an SQLite database (stored as a file on disk).
	db, err := sql.Open("sqlite", "file:tuhuebot.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table.
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        age INTEGER
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	// Insert data into the table.
	insertUserSQL := `INSERT INTO users (name, age) VALUES (?, ?);`

	_, err = db.Exec(insertUserSQL, "Bob", 25)
	if err != nil {
		log.Fatal("Failed to insert user:", err)
	}

	// Query data from the table.
	query := `SELECT id, name, age FROM users;`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Failed to query users:", err)
	}
	defer rows.Close()

	fmt.Println("Users in the database:")
	for rows.Next() {
		var id int
		var name string
		var age int

		// Scan each row into the variables.
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("Failed to scan row:", err)
		}

		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	// Check for errors encountered during iteration.
	if err = rows.Err(); err != nil {
		log.Fatal("Row iteration error:", err)
	}
}
