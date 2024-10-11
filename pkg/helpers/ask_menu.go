package helpers

import (
	tele "gopkg.in/telebot.v3"
)

// Define five inline buttons
var AskMenu = tele.InlineButton{
	Unique: "btn_callback1_ask_menu",
	Text:   "Trở về menu chính",
	Data:   "button1_clicked",
}

// Export the inline keyboard with 5 buttons
var AskMenu_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{AskMenu},
	},
}
