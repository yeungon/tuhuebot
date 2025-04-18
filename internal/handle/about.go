package handle

import (
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

var (
	intro = "Đây là phần mềm dạng bot hỗ trợ tự động, giúp sinh viên tiếp cận thông tin học tập nhanh và thuận lợi hơn. Thông tin mang giá trị tham khảo. \n\nHãy liên hệ trực tiếp với giảng viên hoặc cố vấn học tập để có được trợ giúp chi tiết hơn! \n\n Địa chỉ của bot tại https://t.me/tuhuebot.\n\n Bot hoạt động từ 11/2024 đến nay."

	photo = &tele.Photo{
		Caption: "Mã QR truy cập bot", // Set caption here
		File:    tele.FromURL("https://res.cloudinary.com/yeungon/image/upload/v1728269711/905e81fe3c8985d7dc98_huwerc.jpg"),
	}
)

func dispath(c tele.Context) {
	c.Send(intro)
	c.Send(photo)
	c.Send("⭕🅿♌⚡🌴⛎♈📈❌✌Ⓩ", helpers.AskMenu_InlineKeys)
}

func About(b *tele.Bot) {
	b.Handle("/about", func(c tele.Context) error {
		dispath(c)
		return nil
	})

	b.Handle("about", func(c tele.Context) error {
		dispath(c)
		return nil
	})

	b.Handle(&helpers.Intro, func(c tele.Context) error {
		dispath(c)
		return nil
	})
}
