package pg

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

func GetTimeTable(db *bun.DB) []TimeTable {
	var ctx = context.Background()
	cacheKey := "timetable_data"

	// Check if data is in the cache
	cachedData, err := Cache.Get(cacheKey)
	if err == nil {
		// Cache hit - unmarshal and return cached data
		var timeTable []TimeTable
		err = json.Unmarshal(cachedData, &timeTable)
		if err == nil {
			fmt.Println("Returning data from cache")
			return timeTable
		}
	}

	// Cache miss - query the database
	var timeTable []TimeTable
	err = db.NewSelect().
		Model(&timeTable).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No timetable found.")
			return []TimeTable{} // Return a zero-value slice if no events are found.
		}

		log.Fatal("Failed to retrieve timetable:", err)
	}

	// Store the fetched data in the cache
	data, err := json.Marshal(timeTable)
	if err != nil {
		log.Printf("Failed to marshal timetable data: %v", err)
	} else {
		err = Cache.Set(cacheKey, data)
		if err != nil {
			log.Printf("Failed to cache the result: %v", err)
		}
	}

	fmt.Println("Succeeded fetching data from time_table table stored at XATA.io")
	return timeTable
}
