package users

import (
	"context"
	"log"

	"github.com/uptrace/bun"
)

func GetAllUser(db *bun.DB) []User {
	var ctx = context.Background()
	// Retrieve all users.
	var userList []User
	err := db.NewSelect().
		Model(&userList).
		Order("id ASC").
		Scan(ctx)
	if err != nil {
		log.Fatal("Failed to retrieve users:", err)
	}

	return userList
}

func GetTotalUser(db *bun.DB) int {
	var ctx = context.Background()	
	count, err := db.NewSelect().Model((*User)(nil)).Count(ctx)
	if err != nil {
		log.Fatal("Failed to retrieve total users:", err)
	}
	return count
}
