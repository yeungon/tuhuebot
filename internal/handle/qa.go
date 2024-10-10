package handle

import (
	"fmt"

	"github.com/yeungon/tuhuebot/internal/database"
	"github.com/yeungon/tuhuebot/pkg/cache"
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

		for number := 0; number < 10; number++ {
			fmt.Println("chạy vòng :", number)
			cache.Cache()
		}

		c.Send("Các câu hỏi thường gặp")
		response := database.Query()
		fmt.Println(response)
		for _, record := range response.Records {
			fmt.Printf("ID: %s\n", record.ID)
			fmt.Printf("Question: %s\n", record.Question)
			fmt.Printf("Answer: %s\n\n", record.Answer)
			c.Send("Câu hỏi: " + record.Question)
			c.Send("<strong>Trả lời tham khảo:</strong> " + record.Answer)
		}
		return nil
	})
}
