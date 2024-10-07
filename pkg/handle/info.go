package handle

import (
	tele "gopkg.in/telebot.v3"
)

func Info(b *tele.Bot) {

	// // First message
	// err := c.Send("First message text")
	// if err != nil {
	// 	return err
	// }

	// // Second message
	// err = c.Send("Second message text")
	// if err != nil {
	// 	return err
	// }

	// return nil

	// Define the first inline button
	sodotruong := tele.InlineButton{
		Unique: "btn_callback_sodotruong",
		Text:   "Sơ đồ trường",
		Data:   "button1_clicked",
	}

	// Define the second inline button
	btn2 := tele.InlineButton{
		Unique: "btn_callback2",
		Text:   "Đề cương",
		Data:   "button2_clicked",
	}

	// Create the reply markup and add both buttons in a single row
	inlineKeys := &tele.ReplyMarkup{}
	inlineKeys.InlineKeyboard = [][]tele.InlineButton{
		{sodotruong}, // Row 1: Button 1
		{btn2},       // Row 2: Button 2
	}

	b.Handle("/info", func(c tele.Context) error {

		// Create the reply markup and add the button
		return c.Send("Một số thông tin hữu tích:\n\n", inlineKeys)

	})

	b.Handle(&sodotruong, func(c tele.Context) error {
		return c.Send("Sơ đồ trường!")
	})

	b.Handle(&btn2, func(c tele.Context) error {
		return c.Send("Bạn đang xem các đề thi năm trước, giúp hỗ trợ bạn ôn thi tốt hơn!")
	})

}
