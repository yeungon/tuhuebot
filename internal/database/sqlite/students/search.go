package students

import (
	"context"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

func SearchStudent(db *bun.DB, searchQuery string) []StudentFTS {
	var ctx = context.Background()
	var studentSearch []StudentFTS

	// Quote the search query for exact matching if it contains special characters
	quotedQuery := fmt.Sprintf("\"%s\"", searchQuery)

	err := db.NewSelect().
		Model(&studentSearch).
		Where("students_fts MATCH ?", quotedQuery).
		Limit(100).
		Scan(ctx)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	return studentSearch
}
