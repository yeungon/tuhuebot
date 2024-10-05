package handle

import (
	tele "gopkg.in/telebot.v3"
)

func Qa(b *tele.Bot) {
	b.Handle("/qa", func(c tele.Context) error {
		return c.Reply("hỏi đáp")
	})
}
