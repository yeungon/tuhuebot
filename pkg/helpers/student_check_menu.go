package helpers

import (
	tele "gopkg.in/telebot.v3"
)

// Define five inline buttons
var Back_To_Main_Menu_Second = tele.InlineButton{
	Unique: "btn_callback1_main_menu",
	Text:   "Menu chính ✅",
	Data:   "button1_clicked",
}

// Define five inline buttons
var Back_To_Main_Menu_Assistant = tele.InlineButton{
	Unique: "btn_callback1_main_menu_assistant",
	Text:   "Menu trợ lý ✅",
	Data:   "button1_clicked",
}

// Define five inline buttons
var Keep_Checking_Student = tele.InlineButton{
	Unique: "btn_callback1_keep_check_student",
	Text:   "🔐 Xem sinh viên khác",
	Data:   "button1_clicked",
}

// Export the inline keyboard with 5 buttons
var Student_Check_Menu_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{Back_To_Main_Menu_Second},
		{Back_To_Main_Menu_Assistant, Keep_Checking_Student},
	},
}
