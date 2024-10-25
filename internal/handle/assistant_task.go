package handle

import (
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func AssistantTask(b *tele.Bot) {
	b.Handle(&helpers.Assistant_Tasks, func(c tele.Context) error {
		c.Send("Nội dung này đang cập nhật")
		return nil
	})
}
