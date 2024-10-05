package handle

import (
	tele "gopkg.in/telebot.v3"
)

func Start(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hi, chào mừng bạn đến với bot trợ lý học tập KTU. Chúc bạn một ngày tốt lành. Xin lựa chọn tác vụ phía dưới")
	})
}
