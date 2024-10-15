package users

import (
	"fmt"

	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// Insert a new user into the users table.
func insertUser(db *sql.DB, name string, age int) error {
	insertUserSQL := `INSERT INTO users (name, age) VALUES (?, ?);`
	_, err := db.Exec(insertUserSQL, name, age)
	return err
}

// Retrieve all users from the users table.
func getUsers(db *sql.DB) ([]User, error) {
	query := `SELECT id, name, age FROM users;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

// Define a struct to represent a User.
type User struct {
	ID   int
	Name string
	Age  int
}

func Get(db *sql.DB) {
	// Insert some users.
	if err := insertUser(db, "Alice", 30); err != nil {
		log.Fatal("Failed to insert user:", err)
	}
	if err := insertUser(db, "Bob", 25); err != nil {
		log.Fatal("Failed to insert user:", err)
	}

	// Retrieve and print all users.
	users, err := getUsers(db)
	if err != nil {
		log.Fatal("Failed to retrieve users:", err)
	}

	fmt.Println("Users in the database:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}
