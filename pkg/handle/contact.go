package handle

import (
	tele "gopkg.in/telebot.v3"
)

func Contact(b *tele.Bot) {
	b.Handle("/contact", func(c tele.Context) error {
		return c.Send("Danh bạ liên hệ")
	})
}
