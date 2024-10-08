package handle

import (
	tele "gopkg.in/telebot.v3"
)

func Radio(b *tele.Bot) {
	b.Handle("/radio", func(c tele.Context) error {
		// Send an MP3 file from a URL
		mp3 := &tele.Audio{File: tele.FromURL("https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3")}
		return c.Send(mp3)
	})
}
