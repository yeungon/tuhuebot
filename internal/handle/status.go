package handle

import (
	"fmt"

	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/users"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func FetchStatus(c tele.Context) {
	if helpers.IsAdmin(c) == false {
		c.Send("Bot không xử lý command bạn vừa gửi lên!")
		return
	}
	db := sqlite.DB()
	total := users.GetTotalUser(db)
	info := fmt.Sprintf("Tổng số người dùng hiện tại: %d.", total)
	c.Send(info)
}

func Status(b *tele.Bot) {
	b.Handle("status", func(c tele.Context) error {
		FetchStatus(c)
		return nil
	})
	b.Handle("Status", func(c tele.Context) error {
		FetchStatus(c)
		return nil
	})
	b.Handle("/status", func(c tele.Context) error {
		FetchStatus(c)
		return nil
	})
}
