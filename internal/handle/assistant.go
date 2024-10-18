package handle

import (
	"fmt"

	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func Assistant(b *tele.Bot) {
	content := "Đây là góc hỗ trợ trợ lý. Bạn cần có mật khẩu để truy xuất dữ liệu! Bạn chỉ cần nhập khẩu đúng một lần duy nhất.\n\nXin nhập mật khẩu: "

	b.Handle("/assistant", func(c tele.Context) error {
		return c.Send(content)
	})

	b.Handle(&helpers.Assistant, func(c tele.Context) error {
		return c.Send(content)
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.
		text := c.Text()
		user := c.Sender()
		fmt.Println(user)
		fmt.Println(text)
		return c.Send("Đây là trong Assistant")
	})

}
