package cache

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/allegro/bigcache/v3"
)

var cache *bigcache.BigCache

var random_number = rand.Intn(100)

func Cache() {
	cache, _ = bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

	random_string := strconv.Itoa(random_number)

	fmt.Println("random number inside cache", random_string)

	cache.Set("my-unique-key", []byte(random_string))

	entry, _ := cache.Get("my-unique-key")
	fmt.Println(string(entry))
}
