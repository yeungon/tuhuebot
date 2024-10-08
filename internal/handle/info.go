package handle

import (
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func Info(b *tele.Bot) {

	kehoachnamhoc := tele.InlineButton{
		Unique: "btn_callback_kehoachnamhoc",
		Text:   "Kế hoạch năm học",
		Data:   "button1_clicked",
	}

	sodotruong := tele.InlineButton{
		Unique: "btn_callback_sodotruong",
		Text:   "Sơ đồ trường",
		Data:   "button1_clicked",
	}

	// Define the second inline button
	sotay_sinhvien := tele.InlineButton{
		Unique: "btn_callback2_sotaysinhvien",
		Text:   "Sổ tay sinh viên",
		Data:   "button2_clicked",
	}

	// Create the reply markup and add both buttons in a single row
	inlineKeys := &tele.ReplyMarkup{}
	inlineKeys.InlineKeyboard = [][]tele.InlineButton{
		{sodotruong},     // Row 1: Button 1
		{sotay_sinhvien}, // Row 2: Button 2
		{kehoachnamhoc},
	}

	b.Handle("/info", func(c tele.Context) error {
		// Create the reply markup and add the button
		return c.Send("Một số thông tin hữu tích:\n\n", inlineKeys)

	})

	b.Handle("info", func(c tele.Context) error {
		// Create the reply markup and add the button
		return c.Send("Một số thông tin hữu tích:\n\n", inlineKeys)

	})

	b.Handle(&helpers.Info, func(c tele.Context) error {
		return c.Send("Một số thông tin hữu tích:\n\n", inlineKeys)
	})

	b.Handle(&sodotruong, func(c tele.Context) error {
		link := "https://tieuhoc.org/map/sodo.jpg"
		photo := &tele.Photo{File: tele.FromURL(link)}
		return c.Send(photo)

	})

	b.Handle(&sotay_sinhvien, func(c tele.Context) error {
		link := "https://tieuhoc.org/vanban/quydinh/SOTAYSINHVIEN_2021_tieuhoc.pdf"
		return c.Send(link)
	})

	b.Handle(&kehoachnamhoc, func(c tele.Context) error {
		link := "https://tieuhoc.org/master/2024_2025.jpg"
		photo := &tele.Photo{File: tele.FromURL(link)}
		return c.Send(photo)
	})

}
