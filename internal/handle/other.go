package handle

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

var x = 0

func Other(b *tele.Bot) {
	b.Handle(tele.OnText, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.
		x++
		text := c.Text()
		user := c.Sender()
		fmt.Println(user)
		fmt.Printf("Inside Other %s - %d: ", text, x)
		return c.Send("Sorry, bot không xử lý các thông tin bạn gửi lên!")
	})

	b.Handle(tele.OnPhoto, func(c tele.Context) error {
		fmt.Println(tele.OnPhoto)
		user := c.Sender()
		fmt.Println(user)
		return c.Send("Sorry, bot không xử lý file ảnh bạn gửi lên!")
	})

	b.Handle(tele.OnPoll, func(c tele.Context) error {
		fmt.Println(tele.OnPoll)
		user := c.Sender()
		fmt.Println(user)
		return c.Send("Sorry, bot không xử lý bảng poll bạn gửi lên!")
	})
}
