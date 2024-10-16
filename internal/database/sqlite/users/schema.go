package users

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/bun"
)

// // User represents a user model with Bun ORM.
// type User struct {
// 	ID   int64  `bun:",pk,autoincrement"`
// 	Name string `bun:",notnull"`
// 	Age  int
// }

// User represents a Telegram user or bot with Bun ORM.
type User struct {
	bun.BaseModel `bun:"table:users"` // Specifies the table name for the Bun ORM.

	ID                      int64     `bun:",pk,autoincrement"`      // Primary key with auto-increment.
	TelegramUserID          int64     `bun:",notnull"`               // Telegram user ID.
	IsBot                   bool      `bun:",notnull"`               // Indicates if this user is a bot.
	FirstName               string    `bun:",notnull"`               // User's or bot's first name.
	LastName                *string   `bun:",nullzero"`              // Optional. User's or bot's last name.
	Username                *string   `bun:",nullzero"`              // Optional. User's or bot's username.
	LanguageCode            *string   `bun:",nullzero"`              // Optional. User's language.
	IsPremium               bool      `bun:",notnull,default:false"` // Optional. True if the user is a Premium user.
	AddedToAttachmentMenu   bool      `bun:",notnull,default:false"` // Indicates if the user added the bot to the attachment menu.
	CanJoinGroups           bool      `bun:",notnull,default:false"` // Indicates if the bot can be invited to groups.
	CanReadAllGroupMessages bool      `bun:",notnull,default:false"` // Indicates if privacy mode is disabled.
	SupportsInlineQueries   bool      `bun:",notnull,default:false"` // Indicates if the bot supports inline queries.
	LastSeen                time.Time `bun:",nullzero"`              // Optional. Tracks the last interaction time.
}

func CreateTable(db *bun.DB) {
	var ctx = context.Background()
	// Migrate the schema: create the "users" table if it doesn't exist.
	res, err := db.NewCreateTable().Model((*User)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	info := fmt.Sprintf("The table 'users' created. Rows affected: %d", res)
	fmt.Println(info)
}
