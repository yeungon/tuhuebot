package users

import (
	"context"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

func CreateUser(db *bun.DB) {
	// Insert new users.
	ctx := context.Background()
	users := []*User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	for _, user := range users {
		_, err := db.NewInsert().Model(user).Exec(ctx)
		if err != nil {
			log.Fatal("Failed to insert user:", err)
		}
	}

	fmt.Println("Create new users")
}
