package handle

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yeungon/tuhuebot/internal/database/pg"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

var currentMonth int

func HandleEvent(c tele.Context, month int) {
	pgdata := pg.PG()
	events := pg.GetEvent(pgdata)
	for _, event := range events {
		if int(event.Month) == int(currentMonth+month) {
			c.Send(event.EventData)
		}
	}
	var currentMonthEvents = "Theo dõi các sự kiện: "
	c.Send(currentMonthEvents, helpers.MonthMenu_InlineKeys)
}

func Event(b *tele.Bot) {
	timeLoc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	currentTime := time.Now().In(timeLoc)

	currentMonth = int(currentTime.Month())
	previousMonth := currentMonth - 1
	nextMonth := currentMonth + 1
	// Print the current month
	fmt.Println("Current month:", currentMonth)
	introduction := "Các sự kiện đáng chú ý của tháng " + strconv.Itoa(currentMonth) + ": "

	b.Handle("/event", func(c tele.Context) error {
		c.Send(introduction)
		HandleEvent(c, 0)
		return nil
	})

	b.Handle(&helpers.Event, func(c tele.Context) error {
		c.Send(introduction)
		HandleEvent(c, 0)
		return nil
	})

	b.Handle(&helpers.PreviousMonth, func(c tele.Context) error {
		var currentMonthEvents = "Các sự kiện tháng " + strconv.Itoa(previousMonth)
		c.Send(currentMonthEvents)
		HandleEvent(c, -1)
		return nil
	})

	b.Handle(&helpers.NextMonth, func(c tele.Context) error {
		currentMonthEvents := "Các các sự kiện tháng " + strconv.Itoa(nextMonth)
		c.Send(currentMonthEvents)
		HandleEvent(c, 1)
		return nil
	})
}
