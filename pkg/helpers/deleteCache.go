package helpers

import (
	"fmt"
	"log"

	"github.com/allegro/bigcache/v3"
)

// DeleteCache removes an entry from the cache using the provided key.
func DeleteCache(cache *bigcache.BigCache, key string) {
	err := cache.Delete(key)
	if err != nil {
		if err == bigcache.ErrEntryNotFound {
			fmt.Printf("Cache entry with key '%s' does not exist.\n", key)
		} else {
			log.Printf("Failed to delete cache entry with key '%s': %v\n", key, err)
		}
	} else {
		fmt.Printf("Cache entry with key '%s' has been deleted.\n", key)
	}
}
