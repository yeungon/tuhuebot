package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	Token         string
	AdminID       string
	XATA_API_KEY  string
	XATA_BASE_URL string
}

var once sync.Once
var env *Env

func New() {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		token := os.Getenv("TOKEN")
		admin_id := os.Getenv("ADMIN_ID")
		xata_api_key := os.Getenv("XATA_API_KEY")
		xata_base_url := os.Getenv("XATA_BASE_URL")

		env = &Env{
			Token:         token,
			AdminID:       admin_id,
			XATA_API_KEY:  xata_api_key,
			XATA_BASE_URL: xata_base_url,
		}

	})
}

func Get() *Env {
	return env
}
