package students

import (
	"context"
	"log"

	"github.com/uptrace/bun"
)

func SearchStudent(db *bun.DB, search_querys string) []StudentFTS {
	var ctx = context.Background()
	var student_search []StudentFTS
	err := db.NewSelect().
		Model(&student_search).
		Where("students_fts MATCH ?", search_querys).
		Scan(ctx)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	// for _, student := range student_search {
	// 	fmt.Println("Name:", student.Name)
	// }
	return student_search
}
