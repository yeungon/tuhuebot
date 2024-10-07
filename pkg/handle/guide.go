package handle

import (
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func Guide(b *tele.Bot) {
	content := "Chỉ dẫn giúp bạn học tập tốt hơn!"

	b.Handle("/guide", func(c tele.Context) error {
		return c.Send(content)
	})

	b.Handle(&helpers.Guide, func(c tele.Context) error {
		return c.Send(content)
	})
}
