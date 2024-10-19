package users

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/bun"
)

// SetUserState updates the state of a user based on the provided Telegram ID and state.
func SetUserState(db *bun.DB, telegram_id int64, state bool) {
	var ctx = context.Background()
	var currentUser User
	currentUser.State = state
	currentUser.UpdatedAt = time.Now()
	// Execute the update query to change the state.
	_, err := db.NewUpdate().
		Model(&currentUser).
		Column("state").
		Where("telegram_user_id = ?", telegram_id).
		Exec(ctx)
	if err != nil {
		log.Fatal("There is something wrong with updating the user state: ", err)
	}

	fmt.Printf("Updating state to %v for user with Telegram ID %d\n", state, telegram_id)
}
