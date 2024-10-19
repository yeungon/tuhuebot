package helpers

import (
	tele "gopkg.in/telebot.v3"
)

// Define five inline buttons
var Back_To_Main_Menu = tele.InlineButton{
	Unique: "btn_callback1_main_menu",
	Text:   "Menu chính ⏪",
	Data:   "button1_clicked",
}

// Define five inline buttons
var Keep_Typing_Passcode = tele.InlineButton{
	Unique: "btn_callback1_qa_forward",
	Text:   "⏩ Nhập mật khẩu",
	Data:   "button1_clicked",
}

// Export the inline keyboard with 5 buttons
var Passcode_Menu_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{Back_To_Main_Menu, Keep_Typing_Passcode},
	},
}
