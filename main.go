package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	fmt.Println("Hi, your bot is running as expected!")

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		fmt.Println(c.Text())
		return c.Send("Hi, chào bạn đến với bot học tập KTU. Chúc bạn một ngày tốt lành. Xin lựa chọn tác vụ phía dưới")
	})

	b.Start()
}
