package handle

import (
	"fmt"
	"strconv"

	"github.com/yeungon/tuhuebot/internal/config"
	tele "gopkg.in/telebot.v3"
)

func RemoveCache(b *tele.Bot) {
	content := "Cache từ bigCache đã được xóa!"
	b.Handle("/update", func(c tele.Context) error {
		current_user_id := strconv.FormatInt(c.Sender().ID, 10)
		AdminID := config.Get().AdminID
		if current_user_id != AdminID {
			fmt.Println("Don't handle, the current user is not admin")
			return nil
		}

		// only handle if triggered/called by admin
		fmt.Println(current_user_id)

		return c.Send(content)
	})

}
