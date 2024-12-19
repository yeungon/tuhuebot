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

// SetUserState updates the state of a user based on the provided Telegram ID and state.
func SetUserStateAsking(db *bun.DB, telegram_id int64, state bool) {
	var ctx = context.Background()
	var currentUser User
	currentUser.StateAsking = state
	currentUser.UpdatedAt = time.Now()
	// Execute the update query to change the state_asking.
	_, err := db.NewUpdate().
		Model(&currentUser).
		Column("state_asking").
		Where("telegram_user_id = ?", telegram_id).
		Exec(ctx)
	if err != nil {
		log.Fatal("There is something wrong with updating the user state_asking: ", err)
	}

	fmt.Printf("Updating state to %v for user with Telegram ID %d\n", state, telegram_id)
}

// SetUserState updates the state of a user based on the provided Telegram ID and state.
func SetUserStateChecking(db *bun.DB, telegram_id int64, state bool) {
	var ctx = context.Background()
	var currentUser User
	currentUser.StateChecking = state
	currentUser.UpdatedAt = time.Now()
	// Execute the update query to change the state_asking.
	_, err := db.NewUpdate().
		Model(&currentUser).
		Column("state_checking").
		Where("telegram_user_id = ?", telegram_id).
		Exec(ctx)
	if err != nil {
		log.Fatal("There is something wrong with updating the user state: ", err)
	}

	fmt.Printf("Updating state_checking to %v for user with Telegram ID %d\n", state, telegram_id)
}

func SetUserStateFetching(db *bun.DB, telegram_id int64, state bool) {
	var ctx = context.Background()
	var currentUser User
	currentUser.StateFetching = state
	currentUser.UpdatedAt = time.Now()
	// Execute the update query to change the state_asking.
	_, err := db.NewUpdate().
		Model(&currentUser).
		Column("state_fetching").
		Where("telegram_user_id = ?", telegram_id).
		Exec(ctx)
	if err != nil {
		log.Fatal("There is something wrong with updating the user state: ", err)
	}

	fmt.Printf("Updating state_fetching to %v for user with Telegram ID %d\n", state, telegram_id)
}

// SetUserState updates the state of a user based on the provided Telegram ID and state.
func SetUserLevel(db *bun.DB, telegram_id int64, level int64) {
	var ctx = context.Background()
	var currentUser User
	currentUser.Level = level
	currentUser.UpdatedAt = time.Now()
	// Execute the update query to change the state.
	_, err := db.NewUpdate().
		Model(&currentUser).
		Column("level").
		Where("telegram_user_id = ?", telegram_id).
		Exec(ctx)
	if err != nil {
		log.Fatal("There is something wrong with updating the user state: ", err)
	}

	fmt.Printf("Updating level to %v for user with Telegram ID %d\n", level, telegram_id)
}
