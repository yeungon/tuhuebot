package handle

import (
	"fmt"
	"time"

	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func Event(b *tele.Bot) {
	// Get the current time
	currentTime := time.Now()

	// Extract the month from the current time
	currentMonth := currentTime.Month()

	// Print the current month
	fmt.Println("Current month:", currentMonth)
	content := "Các sự kiện đáng chú ý của tháng " + currentMonth.String() + " :"

	b.Handle("/event", func(c tele.Context) error {
		return c.Send(content)
	})

	b.Handle(&helpers.Event, func(c tele.Context) error {
		return c.Send(content)
	})
}
