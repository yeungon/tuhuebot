package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/yeungon/tuhuebot/internal/config"
	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	"github.com/yeungon/tuhuebot/internal/database/xata"
	tele "gopkg.in/telebot.v3"
)

func Init() {
	config.New()
	xata.Connect()

	db, err := sqlite.Connect()
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	token := config.Get().Token

	Pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(Pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Hi, your bot is running as expected!")
	Handle(b)
	b.Start()
}
