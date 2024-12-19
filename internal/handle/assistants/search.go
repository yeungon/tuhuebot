package assistants

import (
	"fmt"

	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/students"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/users"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func SearchStudent(b *tele.Bot) {
	StartSearchStudent(b)
	KeepSearchingStudent(b)
}

func StartSearchStudent(b *tele.Bot) {
	b.Handle(&helpers.Students_Search, func(c tele.Context) error {
		c.Send("Chế độ tìm kiếm đã bật! 🌿 Nhập thông tin (tên, ngày sinh, quê, hoặc lớp vv...) bất kì để tìm sinh viên! ╰┈➤⤵")
		user_id := c.Sender().ID
		db := sqlite.DB()
		users.SetUserStateFetching(db, user_id, true)
		users.SetUserStateChecking(db, user_id, false)
		c.Send(user_id)
		return nil
	})
}

func KeepSearchingStudent(b *tele.Bot) {
	b.Handle(&helpers.Keep_Searching_Student, func(c tele.Context) error {
		user_id := c.Sender().ID
		db := sqlite.DB()
		users.SetUserStateFetching(db, user_id, true)
		users.SetUserStateChecking(db, user_id, false)
		c.Send("Chế độ tìm kiếm đã được bật lại. Xin nhập từ khóa tìm kiếm! 🌿")
		return nil
	})
}

func StudentSearchFetch(c tele.Context, keyword string) error {
	db := sqlite.DBSTUDENT()
	studentSearch := students.SearchStudent(db, keyword)
	total := len(studentSearch)
	if total == 0 {
		message := fmt.Sprintf("Không tìm thấy thông tin với từ khóa %s.", keyword)
		return c.Send(message)
	}
	result_message := fmt.Sprintf("Tìm được %v kết quả (hiển thị tối đa 100) với từ khóa: %s.", total, keyword)

	c.Send(result_message)
	for _, student := range studentSearch {
		message := fmt.Sprintf(
			"Tên: %s\nMã sinh viên: %s\nGiới tính: %s\nNgày sinh: %s\nLớp: %s\nDân tộc: %s\nCăn cước: %s\nPhone: %s\nEmail: %s\nTỉnh: %s\nĐịa chỉ: %s\nGhi chú: %s",
			student.Name,
			student.StudentCode,
			student.Gender,
			student.DOB,
			student.Class,
			student.Ethnic,
			student.NationalID,
			student.Phone,
			student.Email,
			student.Province,
			student.Address,
			student.Notes,
		)

		if err := c.Send(message); err != nil {
			return fmt.Errorf("failed to send message for student %s: %w", student.Name, err)
		}
	}

	return nil
}
