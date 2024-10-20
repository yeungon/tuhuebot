package users

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

func GetCurrentUser(db *bun.DB, telegram_id int64) User {
	var ctx = context.Background()
	// Retrieve all users.
	var currentUser User
	err := db.NewSelect().
		Model(&currentUser).
		Where("telegram_user_id = ?", telegram_id).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the provided Telegram ID.")
			return User{} // Return a zero-value User if no user is found.
		}

		log.Fatal("Failed to retrieve users:", err)
	}
	return currentUser
}

func GetAllUser(db *bun.DB) []User {
	var ctx = context.Background()
	// Retrieve all users.
	var userList []User
	err := db.NewSelect().
		Model(&userList).
		Order("id ASC").
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found")
			return []User{} // Return a zero-value User if no user is found.
		}
		log.Fatal("Failed to retrieve users:", err)
	}

	return userList
}

func GetTotalUser(db *bun.DB) int {
	var ctx = context.Background()
	count, err := db.NewSelect().Model((*User)(nil)).Count(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0 // Return a zero-value User if no user is found.
		}
		log.Fatal("Failed to retrieve total users:", err)
	}
	return count
}
