package users

import "github.com/uptrace/bun"

func UserState(db *bun.DB, telegram_id int64) bool {
	current_user := GetCurrentUser(db, telegram_id)
	return current_user.State
}

func UserStateAsking(db *bun.DB, telegram_id int64) bool {
	current_user := GetCurrentUser(db, telegram_id)
	return current_user.StateAsking
}
