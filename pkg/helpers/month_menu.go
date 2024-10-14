package helpers

import (
	tele "gopkg.in/telebot.v3"
)

// Define five inline buttons
var PreviousMonth = tele.InlineButton{
	Unique: "btn_callback1_previous_month",
	Text:   "Sự kiện tháng trước ⏪",
	Data:   "button1_clicked",
}

var NextMonth = tele.InlineButton{
	Unique: "btn_callback2_next_month",
	Text:   "⏩ Sự kiện tháng tới",
	Data:   "button2_clicked",
}

// Export the inline keyboard with 5 buttons
var MonthMenu_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{PreviousMonth, NextMonth},
	},
}
