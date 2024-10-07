package helpers

import (
	tele "gopkg.in/telebot.v3"
)

// Define five inline buttons
var Btn1 = tele.InlineButton{
	Unique: "btn_callback1",
	Text:   "Đề thi",
	Data:   "button1_clicked",
}

var Btn2 = tele.InlineButton{
	Unique: "btn_callback2",
	Text:   "Đề cương",
	Data:   "button2_clicked",
}

var Btn3 = tele.InlineButton{
	Unique: "btn_callback3",
	Text:   "Giáo trình",
	Data:   "button3_clicked",
}

var Btn4 = tele.InlineButton{
	Unique: "btn_callback4",
	Text:   "Tài liệu",
	Data:   "button4_clicked",
}

var Info = tele.InlineButton{
	Unique: "thongtinhuuich",
	Text:   "Thông tin hữu ích",
	Data:   "button5_clicked",
}

// Export the inline keyboard with 5 buttons
var MainMenu_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{Btn1}, // Row 1: Button 1
		{Btn2}, // Row 2: Button 2
		{Btn3}, // Row 3: Button 3
		{Btn4}, // Row 4: Button 4
		{Info}, // Row 5: Button 5
	},
}
