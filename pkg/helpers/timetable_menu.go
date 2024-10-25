package helpers

import tele "gopkg.in/telebot.v3"

// Define five inline buttons
var LecturerTimeTable_General = tele.InlineButton{
	Unique: "btn_callback1_timetable_tasks_general",
	Text:   "Thời khóa biểu chung",
	Data:   "button1_clicked",
}

// Define five inline buttons
var LecturerTimeTable_Today = tele.InlineButton{
	Unique: "btn_callback1_timetable_tasks_date_today",
	Text:   "Thời khóa biểu hôm nay",
	Data:   "button1_clicked",
}

// Define five inline buttons
var LecturerTimeTable_Date_Week = tele.InlineButton{
	Unique: "btn_callback1_timetable_tasks_date",
	Text:   "Thời khóa biểu theo ngày",
	Data:   "button1_clicked",
}

// Define five inline buttons
var LecturerTimeTable_Date_Monday = tele.InlineButton{
	Unique: "btn_callback1_timetable_tasks_date_monday",
	Text:   "Thứ Hai",
	Data:   "button1_clicked",
}

// Define five inline buttons
var LecturerTimeTable_Date_Tuesday = tele.InlineButton{
	Unique: "btn_callback1_timetable_tasks_date_tuesday",
	Text:   "Thứ Ba",
	Data:   "button1_clicked",
}

var LecturerTimeTable_Date_Wednesday = tele.InlineButton{
	Unique: "btn_callback1_timetable_tasks_date_wednesday",
	Text:   "Thứ Tư",
	Data:   "button1_clicked",
}

var LecturerTimeTable_Date_Thursday = tele.InlineButton{
	Unique: "btn_callback1_timetable_tasks_date_thursday",
	Text:   "Thứ Năm",
	Data:   "button1_clicked",
}

var LecturerTimeTable_Date_Friday = tele.InlineButton{
	Unique: "btn_callback1_timetable_tasks_date_friday",
	Text:   "Thứ Sáu",
	Data:   "button1_clicked",
}

var LecturerTimeTable_Date_Saturday = tele.InlineButton{
	Unique: "btn_callback1_timetable_tasks_date_saturday",
	Text:   "Thứ Bảy",
	Data:   "button1_clicked",
}

var LecturerTimeTable_Date_Sunday = tele.InlineButton{
	Unique: "btn_callback1_timetable_tasks_date_sunday",
	Text:   "Chủ Nhật",
	Data:   "button1_clicked",
}

var TimeTable_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{LecturerTimeTable_General},
		{LecturerTimeTable_Today},
		{LecturerTimeTable_Date_Week},
	},
}

var TimeTable_InlineKeys_Weekday = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{LecturerTimeTable_Date_Monday}, {LecturerTimeTable_Date_Tuesday},
		{LecturerTimeTable_Date_Wednesday}, {LecturerTimeTable_Date_Thursday},
		{LecturerTimeTable_Date_Friday}, {LecturerTimeTable_Date_Saturday},
		{LecturerTimeTable_Date_Sunday},
	},
}
