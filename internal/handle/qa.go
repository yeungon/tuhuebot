package handle

import (
	"fmt"

	"github.com/yeungon/tuhuebot/internal/database"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func Qa(b *tele.Bot) {
	b.Handle("/qa", func(c tele.Context) error {
		c.Send("Các câu hỏi thường gặp")
		response := database.Query()
		fmt.Println(response)
		return nil
	})

	b.Handle(&helpers.QA, func(c tele.Context) error {
		current_user := c.Sender()

		//return nil

		// for number := 0; number < 10; number++ {
		// 	fmt.Println("chạy vòng :", number)
		// 	//cache.Cache()
		// }

		//c.Send("Các câu hỏi thường gặp")
		response := database.Query()
		//fmt.Println(response)

		for _, record := range response.Records {
			fmt.Printf("ID: %s\n", record.ID)
			questionMsg := "<b>Câu hỏi: </b>" + record.Question
			b.Send(current_user, questionMsg, &tele.SendOptions{
				ParseMode: "HTML",
			})

			answerMsgTexta := "<b>Trả lời: </b>🍟" + record.Answer
			b.Send(current_user, answerMsgTexta, &tele.SendOptions{
				ParseMode: "HTML",
			})
		}

		return nil
	})
}
