package users

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/bun"
)

// User represents a Telegram user or bot with Bun ORM.
type User struct {
	bun.BaseModel           `bun:"table:users"` // Specifies the table name for the Bun ORM.
	ID                      int64               `bun:",pk,autoincrement"`                  // Primary key with auto-increment.
	TelegramUserID          int64               `bun:",notnull,unique"`                    // Telegram user ID.
	IsBot                   bool                `bun:",notnull"`                           // Indicates if this user is a bot.
	Level                   int64               `bun:",notnull,default:1"`                 //
	State                   bool                `bun:",notnull,default:false"`             //
	StateAsking             bool                `bun:",notnull,default:false"`             // If receive the question
	StateBlocking           bool                `bun:",notnull,default:false"`             // If receive the question
	AttemptCounting         int64               `bun:",nullzero,default:0"`                // How many time the user type the password
	QuestionAnswerTracking  int64               `bun:",nullzero,default:0"`                // How many time the user type the password
	ExternalData            int64               `bun:",nullzero,default:0"`                // How many time the user type the password
	FirstName               string              `bun:",nullzero"`                          // User's or bot's first name.
	LastName                *string             `bun:",nullzero"`                          // Optional. User's or bot's last name.
	Username                *string             `bun:",nullzero"`                          // Optional. User's or bot's username.
	LanguageCode            *string             `bun:",nullzero"`                          // Optional. User's language.
	IsPremium               bool                `bun:",notnull,default:false"`             // Optional. True if the user is a Premium user.
	AddedToAttachmentMenu   bool                `bun:",notnull,default:false"`             // Indicates if the user added the bot to the attachment menu.
	CanJoinGroups           bool                `bun:",notnull,default:false"`             // Indicates if the bot can be invited to groups.
	CanReadAllGroupMessages bool                `bun:",notnull,default:false"`             // Indicates if privacy mode is disabled.
	SupportsInlineQueries   bool                `bun:",notnull,default:false"`             // Indicates if the bot supports inline queries.
	CreatedAt               time.Time           `bun:",notnull,default:current_timestamp"` // Created at timestamp
	UpdatedAt               time.Time           `bun:",notnull,default:current_timestamp"` // Updated at timestamp
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
