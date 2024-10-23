package handle

import (
	"github.com/yeungon/tuhuebot/pkg/helpers"
	logging "github.com/yeungon/tuhuebot/pkg/log"
	tele "gopkg.in/telebot.v3"
)

func Log(b *tele.Bot) {
	b.Handle("log", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		file_location := "tuhuebot.json"
		logging.Show(c, file_location)

		return nil
	})

	b.Handle("Log", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		file_location := "tuhuebot.json"
		logging.Show(c, file_location)

		return nil
	})

}
