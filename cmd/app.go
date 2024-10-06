package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func Init() {
	var err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	Pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	// https://github.com/go-telebot/telebot?tab=readme-ov-file#keyboards
	b, err := tele.NewBot(Pref)

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Hi, your bot is running as expected!")
	Handle(b)
	b.Start()
}
