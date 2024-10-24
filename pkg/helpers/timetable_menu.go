package helpers

import tele "gopkg.in/telebot.v3"

// Define five inline buttons
var LecturerTimeTable_General = tele.InlineButton{
	Unique: "btn_callback1_timetable_tasks_general",
	Text:   "Thời khóa biểu chung",
	Data:   "button1_clicked",
}

// Define five inline buttons
var LecturerTimeTable_Date = tele.InlineButton{
	Unique: "btn_callback1_timetable_tasks_date",
	Text:   "Thời khóa biểu theo ngày",
	Data:   "button1_clicked",
}

var TimeTable_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{LecturerTimeTable_General},
		{LecturerTimeTable_Date},
	},
}
