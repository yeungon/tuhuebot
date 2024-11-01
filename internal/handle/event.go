package handle

import (
	"strconv"

	"github.com/yeungon/tuhuebot/internal/database/pg"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func HandleEvent(c tele.Context, month int) {
	pgdata := pg.PG()
	events := pg.GetEvent(pgdata)
	current_thang := helpers.GetCurrentMonth()
	info_thang := int(current_thang + month)

	introduction := "Các sự kiện đáng chú ý của tháng " + strconv.Itoa(info_thang) + ": "
	c.Send(introduction)
	for _, event := range events {
		if int(event.Month) == info_thang {
			c.Send(event.EventData)
		}
	}

	var currentMonthEvents = "⭐⭐⭐Theo dõi các sự kiện: ⭐⭐⭐"
	c.Send(currentMonthEvents, helpers.MonthMenu_InlineKeys)
}

func Event(b *tele.Bot) {
	b.Handle("/event", func(c tele.Context) error {
		HandleEvent(c, 0)
		return nil
	})

	b.Handle(&helpers.Event, func(c tele.Context) error {
		HandleEvent(c, 0)
		return nil
	})

	b.Handle(&helpers.PreviousMonth, func(c tele.Context) error {
		HandleEvent(c, -1)
		return nil
	})

	b.Handle(&helpers.NextMonth, func(c tele.Context) error {
		HandleEvent(c, 1)
		return nil
	})
}
