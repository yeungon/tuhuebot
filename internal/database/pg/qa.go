package pg

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

func GetQuestionAnswer(db *bun.DB) []QA {
	var ctx = context.Background()
	cacheKey := "qa_data"

	// Check if data is in the cache
	cachedData, err := Cache.Get(cacheKey)
	if err == nil {
		// Cache hit - unmarshal and return cached data
		var cachedQAs []QA
		err = json.Unmarshal(cachedData, &cachedQAs)
		if err == nil {
			fmt.Println("Returning data from QA cache named qa_data")
			return cachedQAs
		}
	}

	// Cache miss - query the database
	var question_answer []QA
	err = db.NewSelect().
		Model(&question_answer).
		Order("xata_createdat ASC").
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No question_answer found.")
			return []QA{} // Return a zero-value slice if no events are found.
		}

		log.Fatal("Failed to retrieve question_answer:", err)
	}

	// Store the fetched data in the cache
	data, err := json.Marshal(question_answer)
	if err != nil {
		log.Printf("Failed to marshal question_answer data: %v", err)
	} else {
		err = Cache.Set(cacheKey, data)
		if err != nil {
			log.Printf("Failed to cache the result: %v", err)
		}
	}

	fmt.Println("Succeeded fetching data from qa table stored at XATA.io")
	return question_answer
}

func CreateQA(db *bun.DB, newQA *QA) {
	ctx := context.Background()
	_, err := db.NewInsert().Model(newQA).Exec(ctx)
	if err != nil {
		log.Fatalf("Failed to insert new question: %v", err)
	}

	fmt.Println("New question inserted successfully!")
}
