package handle

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

func Guide(b *tele.Bot) {
	b.Handle("/guide", func(c tele.Context) error {
		user := c.Sender()
		fmt.Println(user)

		return c.Send("Hướng dẫn")
	})
}
