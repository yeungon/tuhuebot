package cmd

import (
	"github.com/yeungon/tuhuebot/internal/handle"
	admin "github.com/yeungon/tuhuebot/internal/handle/admin"
	assistants "github.com/yeungon/tuhuebot/internal/handle/assistants"
	tele "gopkg.in/telebot.v3"
)

func Handle(b *tele.Bot) {
	// Admin----------------
	admin.UpdateCache(b)
	admin.Status(b)
	admin.Log(b)
	admin.Backup(b)
	admin.Test(b)

	// Assistants--level 2--------------
	assistants.Assistant(b)
	assistants.AssistantTask(b)
	assistants.Vanban(b)
	assistants.Calendar(b)
	assistants.Studentlist(b)
	assistants.TimeTable(b)
	assistants.StudentCheck(b)
	assistants.SearchStudent(b)

	// Public--level 1 --------------
	handle.About(b)
	handle.Qa(b)
	handle.Info(b)
	handle.Start(b)
	handle.Tailieu(b)
	handle.Photo(b)
	handle.Radio(b)
	handle.Event(b)

	handle.Profile(b)
	handle.Menu(b)

	handle.Submit(b) //Should be put at the end of the list as this handle will receive the post (submit) from user
}
