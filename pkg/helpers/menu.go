package helpers

import (
	tele "gopkg.in/telebot.v3"
)

// Define five inline buttons
var Intro = tele.InlineButton{
	Unique: "btn_callback1_gioithieu",
	Text:   "Giới thiệu",
	Data:   "button1_clicked",
}

var QA = tele.InlineButton{
	Unique: "btn_callback2_qa",
	Text:   "Hỏi - Đáp",
	Data:   "button2_clicked",
}

var Guide = tele.InlineButton{
	Unique: "btn_callback3_guide",
	Text:   "Hướng dẫn",
	Data:   "button3_clicked",
}

var Materials = tele.InlineButton{
	Unique: "btn_callback4_material",
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
		{Intro},     // Row 1: Button 1
		{QA},        // Row 2: Button 2
		{Guide},     // Row 3: Button 3
		{Materials}, // Row 4: Button 4
		{Info},      // Row 5: Button 5
	},
}
