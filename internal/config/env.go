package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	Token    string
	AdminID  string
	PASSCODE string
	PG       string
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
		pass_code := os.Getenv("PASSCODE")
		postgresq := os.Getenv("PG")

		env = &Env{
			Token:    token,
			AdminID:  admin_id,
			PASSCODE: pass_code,
			PG:       postgresq,
		}

	})
}

func Get() *Env {
	return env
}
