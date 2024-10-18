package handle

import (
	tele "gopkg.in/telebot.v3"
)

func ReceiveQuestion(b *tele.Bot) {
	content := "Đặt câu hỏi!"

	b.Handle("/ask", func(c tele.Context) error {
		return c.Send(content)
	})

	// b.Handle(&helpers.Guide, func(c tele.Context) error {
	// 	return c.Send(content)
	// })
}
