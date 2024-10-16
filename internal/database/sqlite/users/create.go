package users

import (
	"context"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

func CreateUser(db *bun.DB, users []*User) {
	ctx := context.Background()

	for _, user := range users {
		// Check if the user with the same TelegramUserID already exists
		existingUser := new(User)
		userExists, err := db.NewSelect().Model(existingUser).
			Where("telegram_user_id = ?", user.TelegramUserID).
			Exists(ctx)

		// Handle errors while querying
		if err != nil {
			log.Printf("Failed to check if user exists: %v", err)
			continue // Skip to the next user if there's an error
		}

		if userExists {
			// If user already exists, log and skip insertion
			log.Printf("User with TelegramUserID %d already exists. Skipping insertion.", user.TelegramUserID)
			continue
		}

		// No user found; proceed to insert the new user
		_, err = db.NewInsert().Model(user).Exec(ctx)
		if err != nil {
			log.Fatalf("Failed to insert user: %v", err)
		}

		fmt.Println("Created new user with TelegramUserID:", user.TelegramUserID)
	}
}
