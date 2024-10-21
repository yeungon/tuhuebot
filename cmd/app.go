package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/yeungon/tuhuebot/internal/config"
	"github.com/yeungon/tuhuebot/internal/database/pg"
	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	tele "gopkg.in/telebot.v3"
)

func Init() {
	config.New()
	// =============sqlite=============
	sqlite.BunConnect()
	db := sqlite.DB()
	defer db.Close()
	//================PG=================
	pg.Connect()
	pgdatabase := pg.PG().DB
	defer pgdatabase.Close()
	//=================================

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
