package pg

import (
	"time"

	"github.com/uptrace/bun"
)

// User represents a Telegram user or bot with Bun ORM.
type Events struct {
	bun.BaseModel `bun:"table:events"` // Specifies the table name for the Bun ORM.
	XataID        string               `bun:",notnull,unique"`                    // Primary key with auto-increment.
	EventData     string               `bun:",nullzero"`                          //
	EventTasks    string               `bun:",nullzero"`                          //
	EventTracking int64                `bun:",nullzero"`                          //
	EventState    bool                 `bun:",nullzero,default:false"`            //
	Month         int64                `bun:",nullzero"`                          //
	XataCreatedat time.Time            `bun:",notnull,default:current_timestamp"` // Created at timestamp
	XataUpdatedat time.Time            `bun:",notnull,default:current_timestamp"` // Updated at timestamp
}

type QA struct {
	bun.BaseModel `bun:"table:qa"` // Specifies the table name for the Bun ORM.
	XataID        string           `bun:",notnull,unique"`                    // Primary key with auto-increment.
	UserAsked     string           `bun:",nullzero"`                          //
	UserAnswered  string           `bun:",nullzero,default:null"`             // Dự trữ
	Question      string           `bun:",nullzero"`                          //
	Answer        *string          `bun:",nullzero"`                          //
	Published     bool             `bun:",notnull,default:false"`             //
	XataCreatedat time.Time        `bun:",notnull,default:current_timestamp"` // Created at timestamp
	XataUpdatedat time.Time        `bun:",notnull,default:current_timestamp"` // Updated at timestamp
}

type TimeTable struct {
	bun.BaseModel `bun:"table:timetable"` // Specifies the table name for the Bun ORM.
	XataID        string                  `bun:",notnull,unique"`                    // Primary key with auto-increment.
	LecturerName  string                  `bun:",nullzero"`                          //
	Year          int                     `bun:",nullzero"`                          //
	Semester      int                     `bun:",notnull,default:1"`                 //
	Monday        string                  `bun:",nullzero"`                          //
	Tuesday       string                  `bun:",nullzero"`                          //
	Wednesday     string                  `bun:",nullzero"`                          //
	Thursday      string                  `bun:",nullzero"`                          //
	Friday        string                  `bun:",nullzero"`                          //
	Saturday      string                  `bun:",nullzero"`                          //
	Sunday        string                  `bun:",nullzero"`                          //
	XataCreatedat time.Time               `bun:",notnull,default:current_timestamp"` // Created at timestamp
	XataUpdatedat time.Time               `bun:",notnull,default:current_timestamp"` // Updated at timestamp
}
