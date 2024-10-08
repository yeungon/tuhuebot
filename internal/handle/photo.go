package handle

import (
	tele "gopkg.in/telebot.v3"
)

func Photo(b *tele.Bot) {
	b.Handle("/photo", func(c tele.Context) error {
		photo := &tele.Photo{File: tele.FromURL("https://i.ibb.co/LQ5KqFJ/bit-ly-ktu-form.png")}
		return c.Send(photo)

	})
	b.Handle("/link", func(c tele.Context) error {
		return c.Send("Check out this link: https://google.com")
	})

}
