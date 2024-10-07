package handle

import (
	"fmt"

	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func Name(firstName, username string) string {
	if len(firstName) == 0 {
		return username
	} else {
		return firstName
	}
}

func Start(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		var (
			user    = c.Sender()
			intro   = "Welcome onboard"
			welcome = "Chào mừng bạn đến với bot hỗ trợ học tập tự động. Chúc bạn một ngày tốt lành.\n\n Dưới đây là các chức năng, thông tin chính hiện có của bot:"
		)
		firstName := user.FirstName
		username := user.Username
		name := Name(firstName, username)
		introduction := fmt.Sprintf("%s %s. %s", intro, name, welcome)
		return c.Send(introduction, helpers.MainMenu_InlineKeys)
	})
}
