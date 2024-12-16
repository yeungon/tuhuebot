package assistants

import (
	"github.com/yeungon/tuhuebot/internal/config"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func Studentlist(b *tele.Bot) {
	studentList := config.Get().STUDENT_LIST
	b.Handle(&helpers.Students_List, func(c tele.Context) error {
		return c.Send(studentList)
	})
}
