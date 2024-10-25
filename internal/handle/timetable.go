package handle

import (
	"fmt"
	"strings"
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
			message := fmt.Sprintf("%s\n\n", tkb.LecturerName)
			if tkb.Monday != "" {
				message += "✅Thứ hai: " + tkb.Monday + "\n"
			}
			if tkb.Tuesday != "" {
				message += "✅Thứ ba: " + tkb.Tuesday + "\n"
			}
			if tkb.Wednesday != "" {
				message += "✅Thứ tư: " + tkb.Wednesday + "\n"
			}
			if tkb.Thursday != "" {
				message += "✅Thứ năm: " + tkb.Thursday + "\n"
			}
			if tkb.Friday != "" {
				message += "✅Thứ sáu: " + tkb.Friday + "\n"
			}
			if tkb.Saturday != "" {
				message += "✅Thứ bảy: " + tkb.Saturday + "\n"
			}
			if tkb.Sunday != "" {
				message += "✅Chủ nhật: " + tkb.Sunday + "\n"
			}
			if tkb.Notes != "" {
				message += "\n🔥Lưu ý: " + tkb.Notes + "\n"
			}
			c.Send(message)
		}
		return nil
	})

	b.Handle(&helpers.LecturerTimeTable_Date, func(c tele.Context) error {
		day_vietnamese := []string{"Chủ Nhật", "Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy"}
		timeLoc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
		today := time.Now().In(timeLoc)

		currentDay := today.Weekday()
		today_string := strings.ToLower(currentDay.String())

		fmt.Println("Hôm nay là ngày " + today_string)
		dayNumber := int(currentDay)
		fmt.Println(dayNumber)
		c.Send("Thời khóa biểu hôm nay " + day_vietnamese[dayNumber])

		pgdata := pg.PG()
		timetable := pg.GetTimeTableByDay(pgdata, today_string)
		for _, tkb := range timetable {
			lecturer := fmt.Sprintf("%s\n\n", tkb.LecturerName)
			message := ""
			if tkb.Monday != "" {
				if dayNumber == 1 {
					message += "✅Thứ hai: " + tkb.Monday + "\n"
				}
			}
			if tkb.Tuesday != "" {
				if dayNumber == 2 {
					message += "✅Thứ ba: " + tkb.Tuesday + "\n"
				}
			}
			if tkb.Wednesday != "" {
				if dayNumber == 3 {
					message += "✅Thứ tư: " + tkb.Wednesday + "\n"

				}
			}
			if tkb.Thursday != "" {
				if dayNumber == 4 {
					message += "✅Thứ năm: " + tkb.Thursday + "\n"
				}

			}
			if tkb.Friday != "" {
				if dayNumber == 5 {
					message += "✅Thứ sáu: " + tkb.Friday + "\n"
				}

			}
			if tkb.Saturday != "" {
				if dayNumber == 6 {
					message += "✅Thứ bảy: " + tkb.Saturday + "\n"
				}

			}
			if tkb.Sunday != "" {
				if dayNumber == 0 {
					message += "✅Chủ nhật: " + tkb.Sunday + "\n"
				}
			}
			if tkb.Notes != "" {
				message += "\n🔥Lưu ý: " + tkb.Notes + "\n"
			}
			if len(message) > 0 {
				message_send := lecturer + message
				c.Send(message_send)
			}
		}

		return nil
	})
}
