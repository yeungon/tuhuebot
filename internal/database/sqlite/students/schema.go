package students

import (
	"context"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

type StudentFTS struct {
	bun.BaseModel `bun:"table:students_fts,alias:s"`

	ID          int    `bun:"id,pk,nullzero"`        // Corresponds to `id UNINDEXED`
	Name        string `bun:"name,nullzero"`         // Corresponds to `name`
	StudentCode string `bun:"student_code,nullzero"` // Corresponds to `student_code`
	Gender      string `bun:"gender,nullzero"`       // Corresponds to `gender`
	DOB         string `bun:"dob,nullzero"`          // Corresponds to `dob`
	DOBFormat   string `bun:"dob_format,nullzero"`   // Corresponds to `dob_format`
	Class       string `bun:"class,nullzero"`        // Corresponds to `class`
	ClassCode   string `bun:"class_code,nullzero"`   // Corresponds to `class_code`
	Ethnic      string `bun:"ethnic,nullzero"`       // Corresponds to `ethnic`
	NationalID  string `bun:"national_id,nullzero"`  // Corresponds to `national_id`
	Phone       string `bun:"phone,nullzero"`        // Corresponds to `phone`
	Email       string `bun:"email,nullzero"`        // Corresponds to `email`
	Province    string `bun:"province,nullzero"`     // Corresponds to `province`
	Address     string `bun:"address,nullzero"`      // Corresponds to `address`
	Notes       string `bun:"notes,nullzero"`        // Corresponds to `notes`
}

func CreateTableStudents(db *bun.DB) {
	var ctx = context.Background()
	// Migrate the schema: create the "StudentFTS" table if it doesn't exist.
	res, err := db.NewCreateTable().Model((*StudentFTS)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	info := fmt.Sprintf("The table 'StudentFTS' created. Rows affected: %d", res)
	fmt.Println(info)
}
