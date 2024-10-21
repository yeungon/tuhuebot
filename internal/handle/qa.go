package handle

import (
	"fmt"
	"strconv"

	"github.com/yeungon/tuhuebot/internal/database/pg"
	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/users"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

var current int = 1

func FetchQAPG(b *tele.Bot, c tele.Context) {
	c.Send("Các câu hỏi thường gặp")
	current_user := c.Sender()
	pgdata := pg.PG()
	question_answer := pg.GetQuestionAnswer(pgdata)
	for index, record := range question_answer {
		if record.Published == true {
			index_string := strconv.Itoa(index + 1)
			questionMsg := "🌓 🅀🅄🄴🅂🅃🄸🄾🄽 <i>" + index_string + ": " + record.Question + "</i>"
			b.Send(current_user, questionMsg, &tele.SendOptions{
				ParseMode: "HTML",
			})
			// answerMsgTexta := "<b>✅ 🄰🄽🅂🅆🄴🅁: </b>" + record.Answer
			answerMsgTexta := "<b>✅ </b>" + *record.Answer
			b.Send(current_user, answerMsgTexta, &tele.SendOptions{
				ParseMode: "HTML",
			})
		}
	}
}

func Qa(b *tele.Bot) {
	current = current + 1
	b.Handle("/qa", func(c tele.Context) error {
		FetchQAPG(b, c)
		c.Send("Xem các câu hỏi khác", helpers.QA_Menu_InlineKeys)
		return nil
	})

	b.Handle(&helpers.QA, func(c tele.Context) error {
		FetchQAPG(b, c)
		c.Send("Xem các câu hỏi khác", helpers.QA_Menu_InlineKeys)
		return nil
	})
	ControlQuestion(b)
	PostQuestion(b)
}

func ControlQuestion(b *tele.Bot) {
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

func PostQuestion(b *tele.Bot) {
	b.Handle(&helpers.Post_QA, func(c tele.Context) error {
		c.Send(current)
		user_id := c.Sender().ID
		user := c.Sender().ID
		db := sqlite.DB()
		users.SetUserStateAsking(db, user_id, true)
		current_user_asking := users.UserStateAsking(db, user)
		fmt.Println(current_user_asking)
		c.Send("Bot đã bật chế độ nhận câu hỏi. Xin đặt câu hỏi! 🔓")
		return nil
	})
}
