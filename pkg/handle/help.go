package handle

import (
	tele "gopkg.in/telebot.v3"
)

func Help(b *tele.Bot) {
	b.Handle("/help", func(c tele.Context) error {
		return c.Send("Các gợi ý giúp bạn sử dụng bot tốt hơn")
	})
}
