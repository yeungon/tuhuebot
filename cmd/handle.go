package cmd

import (
	"github.com/yeungon/tuhuebot/pkg/handle"
	tele "gopkg.in/telebot.v3"
)

func Handle(b *tele.Bot) {
	handle.About(b)
	handle.Qa(b)
	handle.Guide(b)
	handle.Help(b)
	handle.Start(b)
	handle.Tailieu(b)

	b.Handle(tele.OnText, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.
		return c.Send("Sorry, bot không xử lý các thông tin bạn gửi lên!")
	})
}
