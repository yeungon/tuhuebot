package handle

import (
	"fmt"
	"strconv"

	"github.com/yeungon/tuhuebot/internal/config"
	"github.com/yeungon/tuhuebot/internal/database/pg"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func UpdateCache(b *tele.Bot) {
	b.Handle("update", func(c tele.Context) error {
		handleUpdate(c)
		return nil
	})
	b.Handle("Update", func(c tele.Context) error {
		handleUpdate(c)
		return nil
	})
}

func handleUpdate(c tele.Context) error {
	current_user_id := strconv.FormatInt(c.Sender().ID, 10)
	AdminID := config.Get().AdminID
	if current_user_id != AdminID {
		c.Send("Sorry, bot không xử lý các thông tin người dùng gửi lên!")
		fmt.Println("Don't handle as the current user is not admin")
		return nil
	}
	// only handle if triggered/called by admin
	fmt.Println(current_user_id)
	helpers.DeleteCache(pg.Cache, "events_data")
	helpers.DeleteCache(pg.Cache, "qa_data")
	content := "Cache từ bigCache đã được xóa!"
	return c.Send(content)
}
