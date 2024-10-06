package handle

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

func Start(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		var (
			user    = c.Sender()
			intro   = "Xin chào bạn"
			welcome = "Chào mừng bạn đến với bot hỗ trợ học tập tự động. Chúc bạn một ngày tốt lành."
		)

		fmt.Println(user)
		username := user.Username
		introduction := fmt.Sprintf("%s %s. %s", intro, username, welcome)
		return c.Reply(introduction)
	})
}
