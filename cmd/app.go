package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/yeungon/tuhuebot/internal/config"
	tele "gopkg.in/telebot.v3"
)

func Init() {
	config.NewConfig()
	token := config.Get().Token

	Pref := tele.Settings{
		Token:  token,
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
