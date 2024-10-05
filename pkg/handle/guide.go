package handle

import (
	tele "gopkg.in/telebot.v3"
)

func Guide(b *tele.Bot) {
	b.Handle("/qa", func(c tele.Context) error {
		return c.Send("Hướng dẫn")
	})
}
