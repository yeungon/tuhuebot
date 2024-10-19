package users

import "github.com/uptrace/bun"

func UserLevel(db *bun.DB, telegram_id int64) int {
	current_user := GetCurrentUser(db, telegram_id)
	return int(current_user.Level)
}
