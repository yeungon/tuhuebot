package helpers

import (
	tele "gopkg.in/telebot.v3"
)

// Define five inline buttons
var PreviousMonth_Task = tele.InlineButton{
	Unique: "btn_callback1_previous_month_task_assisstant",
	Text:   "Việc tháng trước ⏪",
	Data:   "button1_clicked",
}

var NextMonth_Task = tele.InlineButton{
	Unique: "btn_callback2_next_month_task_assisstant",
	Text:   "⏩ Việc tháng tới",
	Data:   "button2_clicked",
}

// Export the inline keyboard with 5 buttons
var TaskMenu_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{PreviousMonth_Task, NextMonth_Task},
	},
}
