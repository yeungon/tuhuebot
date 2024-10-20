package helpers

import (
	tele "gopkg.in/telebot.v3"
)

// Define five inline buttons
var Back_QA = tele.InlineButton{
	Unique: "btn_callback1_qa_back",
	Text:   "Câu hỏi trước ⏪",
	Data:   "button1_clicked",
}

// Define five inline buttons
var Post_QA = tele.InlineButton{
	Unique: "btn_callback1_qa_post",
	Text:   "Đặt câu hỏi",
	Data:   "button1_clicked",
}

// Define five inline buttons
var Forward_QA = tele.InlineButton{
	Unique: "btn_callback1_qa_forward",
	Text:   "⏩ Câu hỏi kế",
	Data:   "button1_clicked",
}

// Export the inline keyboard with 5 buttons
var QA_Menu_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{Back_QA, Post_QA, Forward_QA},
	},
}
