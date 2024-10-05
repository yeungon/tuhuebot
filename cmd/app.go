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

	b, err := tele.NewBot(Pref)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Hi, your bot is running as expected!")
	Handle(b)
	b.Handle(tele.OnText, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.
		return c.Send("Sorry, bot không xử lý các thông tin bạn gửi lên!")
	})

	b.Start()
}
