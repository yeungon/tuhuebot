package handle

import (
	"github.com/yeungon/tuhuebot/internal/database"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func Qa(b *tele.Bot) {
	b.Handle("/qa", func(c tele.Context) error {
		database.Query()
		return c.Reply("hỏi đáp")
	})

	b.Handle(&helpers.QA, func(c tele.Context) error {
		database.Query()
		return c.Send("hỏi đáp")
	})
}
