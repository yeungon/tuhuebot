package handle

import (
	"fmt"
	"time"

	"github.com/yeungon/tuhuebot/internal/database/pg"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func TimeTable(b *tele.Bot) {
	b.Handle(&helpers.LecturerTimeTable, func(c tele.Context) error {
		c.Send("Lịch dạy giảng viên:", helpers.TimeTable_InlineKeys)
		return nil
	})
	b.Handle(&helpers.LecturerTimeTable_General, func(c tele.Context) error {
		//c.Send("Thời khóa biểu chung:", helpers.TimeTable_InlineKeys)
		pgdata := pg.PG()
		timetable := pg.GetTimeTable(pgdata)

		for _, tkb := range timetable {
			message := fmt.Sprintf("%s\n\n✅Thứ hai: %s\n✅Thứ ba: %s\n✅Thứ tư: %s\n✅Thứ năm: %s\n✅Thứ sáu: %s\n✅Thứ bảy: %s\n", tkb.LecturerName, tkb.Monday, tkb.Tuesday, tkb.Wednesday, tkb.Thursday, tkb.Friday, tkb.Saturday)
			c.Send(message)
		}
		return nil
	})
	b.Handle(&helpers.LecturerTimeTable_Date, func(c tele.Context) error {
		day_vietnamese := []string{"Chủ Nhật", "Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy"}
		today := time.Now()
		currentDay := today.Weekday()
		dayNumber := int(currentDay)
		c.Send("Thời khóa biểu hôm nay, thứ " + day_vietnamese[dayNumber])
		return nil
	})
}
