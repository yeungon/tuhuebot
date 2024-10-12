package database

import (
	"log"

	"github.com/xataio/xata-go/xata"
	"github.com/yeungon/tuhuebot/internal/config"
)

var SearchClient xata.SearchAndFilterClient
var err error

func Connect() {
	var xata_api_key = config.Get().XATA_API_KEY
	var xata_base_url = config.Get().XATA_BASE_URL
	SearchClient, err = xata.NewSearchAndFilterClient(
		xata.WithAPIKey(xata_api_key),
		xata.WithBaseURL(xata_base_url),
	)
	if err != nil {
		log.Fatalf("Failed to create Search and Filter client: %v", err)
	}
}
