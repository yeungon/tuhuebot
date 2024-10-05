package handle

import (
	tele "gopkg.in/telebot.v3"
)

func Tailieu(b *tele.Bot) {
	b.Handle("/qa", func(c tele.Context) error {
		return c.Send("Hi, tài liệu")
	})
}
