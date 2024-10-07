package handle

import (
	tele "gopkg.in/telebot.v3"
)

func About(b *tele.Bot) {
	b.Handle("/about", func(c tele.Context) error {
		intro := "Đây là phần mềm dạng bot hỗ trợ tự động, giúp sinh viên tiếp cận thông tin học tập nhanh và thuận lợi hơn. Thông tin mang giá trị tham khảo. \n\nHãy liên hệ trực tiếp với giảng viên hoặc cố vấn học tập để có được trợ giúp chi tiết hơn!"
		// photo := &tele.Photo{
		// 	Caption: intro, // Set caption here
		// 	File:    tele.FromURL("https://res.cloudinary.com/yeungon/image/upload/v1728269711/905e81fe3c8985d7dc98_huwerc.jpg"),
		// }

		return c.Send(intro)

	})
}
