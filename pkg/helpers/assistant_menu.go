package helpers

import (
	tele "gopkg.in/telebot.v3"
)

// Define five inline buttons
var LecturerTimeTable = tele.InlineButton{
	Unique: "btn_callback1_time_table_intro",
	Text:   "Lịch dạy giảng viên",
	Data:   "button1_clicked",
}

// Define five inline buttons
var GeneralTimeTable = tele.InlineButton{
	Unique: "btn_callback1_time_table_general",
	Text:   "Lịch dạy tổng thể",
	Data:   "button1_clicked",
}

var DetailTimeTable = tele.InlineButton{
	Unique: "btn_callback2_time_table_detail",
	Text:   "Lịch dạy chi tiết",
	Data:   "button2_clicked",
}

// Define five inline buttons
var Assistant_Tasks = tele.InlineButton{
	Unique: "btn_callback1_time_assistant_tasks",
	Text:   "Công việc trợ lý theo tháng",
	Data:   "button1_clicked",
}

// Export the inline keyboard with 5 buttons
var Assitant_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{LecturerTimeTable},
		{Assistant_Tasks},
	},
}
