package pg

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

func GetEvent(db *bun.DB) []Events {
	var ctx = context.Background()
	cacheKey := "events_data"

	// Check if data is in the cache
	cachedData, err := cache.Get(cacheKey)
	if err == nil {
		// Cache hit - unmarshal and return cached data
		var cachedEvents []Events
		err = json.Unmarshal(cachedData, &cachedEvents)
		if err == nil {
			fmt.Println("Returning data from cache")
			return cachedEvents
		}
	}

	// Cache miss - query the database
	var event []Events
	err = db.NewSelect().
		Model(&event).
		Order("month ASC").
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No events found.")
			return []Events{} // Return a zero-value slice if no events are found.
		}

		log.Fatal("Failed to retrieve events:", err)
	}

	// Store the fetched data in the cache
	data, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event data: %v", err)
	} else {
		err = cache.Set(cacheKey, data)
		if err != nil {
			log.Printf("Failed to cache the result: %v", err)
		}
	}

	fmt.Println("Succeeded fetching data from events table stored at XATA.io")
	return event
}
