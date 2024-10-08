package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	Token   string
	AdminID string
}

var once sync.Once
var env *Env

func NewConfig() *Env {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		token := os.Getenv("TOKEN")
		admin_id := os.Getenv("ADMIN_ID")
		env = &Env{
			Token:   token,
			AdminID: admin_id,
		}

	})

	return env

}
