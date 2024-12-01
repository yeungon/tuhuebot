package handle

import (
	"strconv"

	"github.com/yeungon/tuhuebot/internal/database/pg"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func HandleAssistantTask(c tele.Context, month int) {
	pgdata := pg.PG()
	events := pg.GetEvent(pgdata)
	current_thang := helpers.GetCurrentMonth()
	info_thang := int(current_thang + month)

	if info_thang > 12 {
		info_thang = 1
	}

	introduction := "Các công việc cần chú ý của tháng " + strconv.Itoa(info_thang) + ": "
	c.Send(introduction)
	for _, event := range events {
		if int(event.Month) == info_thang {
			c.Send(event.EventTasks)
		}
	}
	var currentMonthEvents = "Theo dõi các nhiệm vụ theo tháng: "
	c.Send(currentMonthEvents, helpers.TaskMenu_InlineKeys)
}

func AssistantTask(b *tele.Bot) {
	b.Handle(&helpers.Assistant_Tasks, func(c tele.Context) error {

		HandleAssistantTask(c, 0)
		return nil
	})

	b.Handle(&helpers.PreviousMonth_Task, func(c tele.Context) error {
		HandleAssistantTask(c, -1)
		return nil
	})

	b.Handle(&helpers.NextMonth_Task, func(c tele.Context) error {
		HandleAssistantTask(c, 1)
		return nil
	})
}
