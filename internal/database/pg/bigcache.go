package pg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
)

// Global variable for BigCache
var Cache *bigcache.BigCache

// Initialize BigCache
func init() {
	var err error
	config := bigcache.Config{
		Shards:     1024,                // Number of cache shards
		LifeWindow: 20000 * time.Minute, // Cache entries will expire after 10 minutes
		Verbose:    true,
		OnRemoveWithReason: func(key string, entry []byte, reason bigcache.RemoveReason) {
			fmt.Printf("Entry with key '%s' was removed! Reason: %v\n", key, reason)
		},
	}
	Cache, err = bigcache.New(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to initialize BigCache: %v", err)
	}
}
