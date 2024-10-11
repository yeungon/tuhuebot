package handle

import (
	"fmt"
	"strconv"

	"github.com/yeungon/tuhuebot/internal/database"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func FetchQa(b *tele.Bot, c tele.Context) {
	current_user := c.Sender()
	c.Send("Các câu hỏi thường gặp")
	response := database.Query()

	for index, record := range response.Records {
		fmt.Printf("ID: %s\n", record.ID)
		index_string := strconv.Itoa(index + 1)
		questionMsg := "<b>🌓 🅀🅄🄴🅂🅃🄸🄾🄽 </b><i>" + index_string + ": " + record.Question + "</i>"
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

func Qa(b *tele.Bot) {
	b.Handle("/qa", func(c tele.Context) error {
		FetchQa(b, c)
		return nil
	})

	b.Handle(&helpers.QA, func(c tele.Context) error {
		FetchQa(b, c)
		return nil
	})
}
