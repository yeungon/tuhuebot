package handle

import (
	"fmt"
	"strconv"

	"github.com/yeungon/tuhuebot/internal/config"
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
		fmt.Println("Don't handle, the current user is not admin")
		return nil
	}
	// only handle if triggered/called by admin
	fmt.Println(current_user_id)
	content := "Cache từ bigCache đã được xóa!"
	return c.Send(content)
}
