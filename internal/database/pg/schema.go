package pg

import (
	"time"

	"github.com/uptrace/bun"
)

// User represents a Telegram user or bot with Bun ORM.
type Events struct {
	bun.BaseModel `bun:"table:events"` // Specifies the table name for the Bun ORM.
	XataID        string               `bun:",notnull,unique"`                    // Primary key with auto-increment.
	EventData     string               `bun:",nullzero"`                          // Indicates if this user is a bot.
	Month         int64                `bun:",nullzero"`                          //
	XataCreatedat time.Time            `bun:",notnull,default:current_timestamp"` // Created at timestamp
	XataUpdatedat time.Time            `bun:",notnull,default:current_timestamp"` // Updated at timestamp
}
