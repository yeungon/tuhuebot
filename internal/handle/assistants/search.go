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
		c.Send("Chế độ tìm kiếm đã bật! Nhập thông tin (tên, ngày sinh, quê, hoặc lớp vv...) bất kì để tìm sinh viên! ╰┈➤⤵")
		user_id := c.Sender().ID
		db := sqlite.DB()
		users.SetUserStateFetching(db, user_id, true)
		fmt.Println(users.UserStateFetching(db, user_id))
		c.Send(user_id)
		return nil
	})
}

func KeepSearchingStudent(b *tele.Bot) {
	b.Handle(&helpers.Keep_Searching_Student, func(c tele.Context) error {
		user_id := c.Sender().ID
		db := sqlite.DB()
		users.SetUserStateFetching(db, user_id, true)
		c.Send("Chế độ tìm kiếm đã được bật lại. Xin nhập từ khóa tìm kiếm! 😀")
		return nil
	})
}

func StudentSearchFetch(c tele.Context, keyword string) error {
	db := sqlite.DBSTUDENT()
	student_search := students.SearchStudent(db, keyword)
	for _, student := range student_search {
		fmt.Println(student.Name)
		c.Send(student.Name)
	}
	return nil
}
