package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
)

type XataData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func TestCache() {
	// Configure the cache with 10 minutes expiration
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

	if err != nil {
		log.Fatal(err)
	}

	// Simulate fetching data from Xata.io
	dataFromXata := XataData{
		Name:  "Sample Item",
		Value: 42,
	}

	// Serialize the data to JSON
	jsonData, err := json.Marshal(dataFromXata)
	if err != nil {
		log.Fatal(err)
	}

	// Store the JSON in BigCache with a key
	key := "xataData"
	err = cache.Set(key, jsonData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data stored in cache.")

	// Retrieve the data from cache
	cachedData, err := cache.Get(key)
	if err != nil {
		log.Fatal(err)
	}

	// Deserialize the JSON back into a struct
	var retrievedData XataData
	err = json.Unmarshal(cachedData, &retrievedData)
	if err != nil {
		log.Fatal(err)
	}

	// Print the retrieved data
	fmt.Printf("Retrieved data: %+v\n", retrievedData)
}
