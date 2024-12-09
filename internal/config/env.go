package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	Token        string
	AdminID      string
	PASSCODE     string
	PG           string
	SPH_USERNAME string
	SPH_PASSWORD string
	STUDENT_LIST string
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
		sph_username := os.Getenv("SPH_USERNAME")
		sph_password := os.Getenv("SPH_PASSWORD")
		student_list := os.Getenv("STUDENT_LIST")

		env = &Env{
			Token:        token,
			AdminID:      admin_id,
			PASSCODE:     pass_code,
			PG:           postgresq,
			SPH_USERNAME: sph_username,
			SPH_PASSWORD: sph_password,
			STUDENT_LIST: student_list,
		}

	})
}

func Get() *Env {
	return env
}
