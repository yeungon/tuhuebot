package handle

import (
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func Tailieu(b *tele.Bot) { // Define the inline button
	// Define the first inline button
	dethi := tele.InlineButton{
		Unique: "btn_callback1",
		Text:   "Đề thi",
		Data:   "button1_clicked",
	}

	// Define the second inline button
	decuong := tele.InlineButton{
		Unique: "btn_callback2",
		Text:   "Đề cương",
		Data:   "button2_clicked",
	}

	// Create the reply markup and add both buttons in a single row
	inlineKeys := &tele.ReplyMarkup{}
	inlineKeys.InlineKeyboard = [][]tele.InlineButton{
		{dethi},   // Row 1: Button 1
		{decuong}, // Row 2: Button 2
	}

	b.Handle("/tailieu", func(c tele.Context) error {
		// Create the reply markup and add the button
		return c.Send("Bạn đang duyệt kho tài liệu học tập:", inlineKeys)

	})

	b.Handle(&helpers.Materials, func(c tele.Context) error {
		// Create the reply markup and add the button
		return c.Send("Bạn đang duyệt kho tài liệu học tập:", inlineKeys)

	})

	b.Handle(&dethi, func(c tele.Context) error {
		return c.Send("Tuyển tập các đề thi năm trước (inside handle/tailieu)!")
	})

	b.Handle(&decuong, func(c tele.Context) error {
		return c.Send("Bạn đang xem các đề thi năm trước, giúp hỗ trợ bạn ôn thi tốt hơn!")
	})

}
