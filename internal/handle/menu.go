package handle

import (
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func Menu(b *tele.Bot) {
	b.Handle(&helpers.AskMenu, func(c tele.Context) error {
		intro := "🅰🅱©↩📧🎏⛽♓ℹ🗾🎋👢Ⓜ♑⭕🅿♌⚡🌴⛎✌Ⓩ"
		return c.Send(intro, helpers.MainMenu_InlineKeys)
	})

}
