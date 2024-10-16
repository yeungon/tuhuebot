package users

import (
	"context"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

// User represents a user model with Bun ORM.
type User struct {
	ID   int64  `bun:",pk,autoincrement"`
	Name string `bun:",notnull"`
	Age  int
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
