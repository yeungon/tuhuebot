package pg

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

func GetEvent(db *bun.DB) []Events {
	var ctx = context.Background()
	// Retrieve all users.
	var event []Events
	err := db.NewSelect().
		Model(&event).
		Order("month ASC").
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the provided Telegram ID.")
			return []Events{} // Return a zero-value User if no user is found.
		}

		log.Fatal("Failed to retrieve users:", err)
	}
	fmt.Println("Fetching event ok")
	return event
}
