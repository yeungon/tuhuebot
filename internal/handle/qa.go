package handle

import (
	"fmt"
	"strconv"

	"github.com/yeungon/tuhuebot/internal/database/qa"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

var current int = 1

func FetchQa(b *tele.Bot, c tele.Context) {
	current_user := c.Sender()
	c.Send("Các câu hỏi thường gặp")
	response := qa.QueryQA()

	for index, record := range response.Records {
		//fmt.Printf("ID: %s\n", record.ID)
		fmt.Println("Published:", record.Published)
		if record.Published == true {
			index_string := strconv.Itoa(index + 1)
			questionMsg := "🌓 🅀🅄🄴🅂🅃🄸🄾🄽 <i>" + index_string + ": " + record.Question + "</i>"
			// Reference emoji icon havested here: https://emojipedia.org/first-quarter-moon
			b.Send(current_user, questionMsg, &tele.SendOptions{
				ParseMode: "HTML",
			})
			// answerMsgTexta := "<b>✅ 🄰🄽🅂🅆🄴🅁: </b>" + record.Answer
			answerMsgTexta := "<b></b>" + record.Answer
			b.Send(current_user, answerMsgTexta, &tele.SendOptions{
				ParseMode: "HTML",
			})
		}
	}
}

func Qa(b *tele.Bot) {
	current = current + 1
	b.Handle("/qa", func(c tele.Context) error {
		FetchQa(b, c)
		c.Send("Xem các câu hỏi khác", helpers.QA_Menu_InlineKeys)
		return nil
	})

	b.Handle(&helpers.QA, func(c tele.Context) error {
		FetchQa(b, c)
		c.Send("Xem các câu hỏi khác", helpers.QA_Menu_InlineKeys)
		return nil
	})

	b.Handle(&helpers.Back_QA, func(c tele.Context) error {
		//FetchQa(b, c)
		//c.Send("Xem các câu hỏi khác", helpers.QA_Menu_InlineKeys)
		current = current - 1
		current := strconv.Itoa(int(current))
		c.Send(current)
		fmt.Println(current)
		c.Send("test back")
		return nil
	})

	b.Handle(&helpers.Forward_QA, func(c tele.Context) error {
		//FetchQa(b, c)
		//c.Send("Xem các câu hỏi khác", helpers.QA_Menu_InlineKeys)
		current = current + 1
		current := strconv.Itoa(int(current))
		c.Send(current)
		fmt.Println(current)
		c.Send("test forward")
		return nil
	})
}
