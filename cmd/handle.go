package cmd

import (
	"github.com/yeungon/tuhuebot/internal/handle"
	tele "gopkg.in/telebot.v3"
)

func Handle(b *tele.Bot) {
	handle.About(b)
	handle.Qa(b)
	handle.Assistant(b)
	handle.Info(b)
	handle.Start(b)
	handle.Tailieu(b)
	handle.Photo(b)
	handle.Radio(b)
	handle.Event(b)
	handle.Test(b)
	handle.Profile(b)
	handle.Menu(b)
	handle.UpdateCache(b)
	handle.TimeTable(b)
	handle.Status(b)
	handle.Log(b)
	handle.Submit(b) //Should be put at the end of the list.
}
