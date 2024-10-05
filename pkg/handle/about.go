package handle

import (
	tele "gopkg.in/telebot.v3"
)

func About(b *tele.Bot) {
	b.Handle("/about", func(c tele.Context) error {
		return c.Send("Đây là phần mềm dạng bot hỗ trợ tự động, giúp sinh viên tiếp cận thông tin học tập nhanh và thuận lợi hơn. Thông tin mang giá trị tham khảo. Hãy liên hệ trực tiếp với giảng viên hướng dẫn hoặc cố vấn học tập để có được trợ giúp chi tiết hơn!")
	})
}
