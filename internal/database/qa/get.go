package qa

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/xataio/xata-go/xata"
	"github.com/yeungon/tuhuebot/internal/database"
)

var cache *bigcache.BigCache

// Initialize BigCache
func init() {
	var err error
	fmt.Print("Running? - inside database/database/QA/get")
	cache, err = bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		log.Fatalf("Failed to initialize BigCache: %v", err)
	}
}

type Response struct {
	Meta    Meta     `json:"meta"`
	Records []Record `json:"records"`
}

type Meta struct {
	Page Page `json:"page"`
}

type Page struct {
	Cursor string `json:"cursor"`
	More   bool   `json:"more"`
	Size   int    `json:"size"`
}

type Record struct {
	ID        string `json:"id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	Xata      Xata   `json:"xata"`
	Published bool   `json:"published"`
}

type Xata struct {
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Version   int    `json:"version"`
}

func QueryQA() Response {
	// Cache key (can be dynamic if needed based on the query)
	cacheKey := "qa_table_data"

	// Check if data is in the cache
	cachedData, err := cache.Get(cacheKey)
	if err == nil {
		// Cache hit - return the cached data
		var cachedResponse Response
		err = json.Unmarshal(cachedData, &cachedResponse)
		if err == nil {
			log.Println("Returning data from cache")
			return cachedResponse
		}
	}

	// Cache miss - Query the "qa" table
	qa, err := database.SearchClient.Query(context.TODO(), xata.QueryTableRequest{
		BranchRequestOptional: xata.BranchRequestOptional{
			DatabaseName: xata.String("tuhuebot"),
			BranchName:   xata.String("main"),
		},
		TableName: "qa",
		// Reference: https://github.com/xataio/xata-go/blob/main/xata/search_filter_client_test.go
		// Payload: xata.QueryTableRequestPayload{
		// 	// Columns: []string{"question", "answer"},
		// 	Filter: &xata.FilterExpression{
		// 		Any: xata.NewFilterListFromFilterExpressionList([]*xata.FilterExpression{
		// 			{
		// 				Exists: xata.String("published"), // Check if 'published' exists (both true and false work)
		// 			},
		// 		}),
		// 	},
		// },
	})
	if err != nil {
		log.Fatalf("Error querying the qa table: %v", err)
	}

	// Convert data to a readable JSON format
	qa_JSON, err := json.MarshalIndent(qa, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal qa response: %v", err)
	}

	var response Response
	err = json.Unmarshal([]byte(qa_JSON), &response)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Store the result in the cache
	err = cache.Set(cacheKey, qa_JSON)
	if err != nil {
		log.Printf("Failed to cache the result: %v", err)
	}

	return response
}
