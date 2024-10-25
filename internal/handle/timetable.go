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
		c.Send("Lịch dạy giảng viên")
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

	b.Handle(&helpers.LecturerTimeTable_Today, func(c tele.Context) error {
		day_vietnamese := []string{"Chủ Nhật", "Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy"}
		timeLoc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
		today := time.Now().In(timeLoc)
		currentDay := today.Weekday()
		today_string := strings.ToLower(currentDay.String())
		dayNumber := int(currentDay)
		c.Send("Thời khóa biểu hôm nay " + day_vietnamese[dayNumber])
		fetchTimeTable(c, today_string, dayNumber)
		return nil
	})

	b.Handle(&helpers.LecturerTimeTable_Date_Week, func(c tele.Context) error {
		c.Send("Lịch dạy giảng viên theo ngày:", helpers.TimeTable_InlineKeys_Weekday)
		return nil
	})

	b.Handle(&helpers.LecturerTimeTable_Date_Monday, func(c tele.Context) error {
		c.Send("Lịch dạy giảng viên Thứ Hai:")
		fetchTimeTable(c, "monday", 1)
		return nil
	})

	b.Handle(&helpers.LecturerTimeTable_Date_Tuesday, func(c tele.Context) error {
		c.Send("Lịch dạy giảng viên Thứ Ba:")
		fetchTimeTable(c, "tuesday", 2)
		return nil
	})

	b.Handle(&helpers.LecturerTimeTable_Date_Wednesday, func(c tele.Context) error {
		c.Send("Lịch dạy giảng viên Thứ Tư:")
		fetchTimeTable(c, "wednesday", 3)
		return nil
	})

	b.Handle(&helpers.LecturerTimeTable_Date_Thursday, func(c tele.Context) error {
		c.Send("Lịch dạy giảng viên Thứ Năm:")
		fetchTimeTable(c, "thursday", 4)
		return nil
	})

	b.Handle(&helpers.LecturerTimeTable_Date_Friday, func(c tele.Context) error {
		c.Send("Lịch dạy giảng viên Thứ Sáu:")
		fetchTimeTable(c, "friday", 5)
		return nil
	})

	b.Handle(&helpers.LecturerTimeTable_Date_Saturday, func(c tele.Context) error {
		c.Send("Lịch dạy giảng viên Thứ Bảy:")
		fetchTimeTable(c, "saturday", 6)
		return nil
	})

	b.Handle(&helpers.LecturerTimeTable_Date_Sunday, func(c tele.Context) error {
		c.Send("Lịch dạy giảng viên Chủ Nhật:")
		fetchTimeTable(c, "sunday", 0)
		return nil
	})

}

func fetchTimeTable(c tele.Context, dayFetching string, dayNumber int) {
	pgdata := pg.PG()
	timetable := pg.GetTimeTableByDay(pgdata, dayFetching)
	for _, tkb := range timetable {
		lecturer := fmt.Sprintf("%s\n", tkb.LecturerName)
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
		if len(message) > 0 {
			if tkb.Notes != "" {
				message += "\n🔥Lưu ý: " + tkb.Notes + "\n"
			}
			message_send := lecturer + message
			c.Send(message_send)
		}
	}

}
