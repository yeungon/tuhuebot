package handle

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yeungon/tuhuebot/internal/database/pg"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func HandleAssistantTask(c tele.Context, month int) {
	pgdata := pg.PG()
	events := pg.GetEvent(pgdata)
	for _, event := range events {
		if int(event.Month) == int(currentMonth+month) {
			c.Send(event.EventTasks)
		}
	}
	var currentMonthEvents = "Theo dõi các nhiệm vụ theo tháng: "
	c.Send(currentMonthEvents, helpers.TaskMenu_InlineKeys)
}

func AssistantTask(b *tele.Bot) {
	timeLoc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	currentTime := time.Now().In(timeLoc)
	currentMonth = int(currentTime.Month())
	previousMonth := currentMonth - 1
	nextMonth := currentMonth + 1

	// Print the current month
	fmt.Println("Current month:", currentMonth)
	introduction := "Các công việc cần chú ý của tháng " + strconv.Itoa(currentMonth) + ": "

	b.Handle(&helpers.Assistant_Tasks, func(c tele.Context) error {
		c.Send(introduction)
		HandleAssistantTask(c, 0)
		return nil
	})

	b.Handle(&helpers.PreviousMonth_Task, func(c tele.Context) error {
		var currentMonthEvents = "Các công việc tháng " + strconv.Itoa(previousMonth)
		c.Send(currentMonthEvents)
		HandleAssistantTask(c, -1)
		return nil
	})

	b.Handle(&helpers.NextMonth_Task, func(c tele.Context) error {
		currentMonthEvents := "Các công việc tháng " + strconv.Itoa(nextMonth)
		c.Send(currentMonthEvents)
		HandleAssistantTask(c, 1)
		return nil
	})
}
