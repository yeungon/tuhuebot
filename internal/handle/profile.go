package handle

import (
	tele "gopkg.in/telebot.v3"
)

func Profile(b *tele.Bot) {
	b.Handle("/profile", func(c tele.Context) error {
		return c.Send("Work in progress - Check out this link: https://google.com")
	})

}
