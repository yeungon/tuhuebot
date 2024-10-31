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
var Assistant_Tasks = tele.InlineButton{
	Unique: "btn_callback1_time_assistant_tasks",
	Text:   "Công việc trợ lý theo tháng",
	Data:   "button1_clicked",
}

// Define five inline buttons
var Tracking_Announcement = tele.InlineButton{
	Unique: "btn_callback1_time_tracking_announcement",
	Text:   "Theo dõi công văn mới nhất",
	Data:   "button1_clicked",
}

var Assitant_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{LecturerTimeTable},
		{Assistant_Tasks},
		{Tracking_Announcement},
	},
}
