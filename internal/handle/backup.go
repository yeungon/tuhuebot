package handle

import (
	tele "gopkg.in/telebot.v3"
)

func Backup(b *tele.Bot) {
	b.Handle("backup", func(c tele.Context) error {
		intro := "🅰🅱©↩📧🎏⛽♓ℹ🗾🎋👢Ⓜ♑⭕🅿♌⚡🌴⛎✌Ⓩ"
		return c.Send(intro, "manual backup")
	})
}
