package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/yeungon/tuhuebot/internal/config"
	"github.com/yeungon/tuhuebot/internal/database/bbolt"
	"github.com/yeungon/tuhuebot/internal/database/pg"
	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	logging "github.com/yeungon/tuhuebot/pkg/log"
	tele "gopkg.in/telebot.v3"
)

func Init() {
	logging.Log()

	config.New()
	// =============sqlite=============
	sqlite.BunConnect()
	db := sqlite.DB()
	defer db.Close()
	//================PG=================
	pg.Connect()
	pgdatabase := pg.PG().DB
	defer pgdatabase.Close()
	//===============bbolt==================
	dbbolt := bbolt.Connect()
	defer dbbolt.Close()

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

	// When shutting down the application, ensure the log is closed properly
	defer logging.CloseLog()
	b.Start()
}
