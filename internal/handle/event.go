package handle

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func Event(b *tele.Bot) {
	// Get the current time
	currentTime := time.Now()

	// Extract the month from the current time
	currentMonth := int(currentTime.Month())
	previousMonth := currentMonth - 1
	nextMonth := currentMonth + 1

	// Print the current month
	fmt.Println("Current month:", currentMonth)
	content := "Các sự kiện đáng chú ý của tháng " + strconv.Itoa(currentMonth) + ": "

	b.Handle("/event", func(c tele.Context) error {
		c.Send(content)
		var currentMonthEvents = "Theo dõi các sự kiện: "
		return c.Send(currentMonthEvents, helpers.MonthMenu_InlineKeys)
	})

	b.Handle(&helpers.Event, func(c tele.Context) error {
		c.Send(content)
		var currentMonthEvents = "Theo dõi các sự kiện: "
		return c.Send(currentMonthEvents, helpers.MonthMenu_InlineKeys)
	})

	b.Handle(&helpers.PreviousMonth, func(c tele.Context) error {
		var currentMonthEvents = "Các sự kiện tháng " + strconv.Itoa(previousMonth)
		return c.Send(currentMonthEvents)
	})

	b.Handle(&helpers.NextMonth, func(c tele.Context) error {
		var currentMonthEvents = "Các các sự kiện tháng " + strconv.Itoa(nextMonth)
		return c.Send(currentMonthEvents)
	})
}
