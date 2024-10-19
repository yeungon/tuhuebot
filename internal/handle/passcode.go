package handle

import (
	"fmt"

	"github.com/yeungon/tuhuebot/internal/config"
	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/users"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func check_passcode(user_input string) bool {
	passcode := config.Get().PASSCODE
	fmt.Println(passcode)
	if passcode == user_input {
		return true
	}
	return false
}

func Passcode(b *tele.Bot) {
	b.Handle(tele.OnText, func(c tele.Context) error {
		user := c.Sender().ID
		db := sqlite.DB()
		state := users.UserState(db, user)
		if state == true {
			user_input := c.Text()
			passcode := check_passcode(user_input)
			users.SetUserState(db, user, false)

			if passcode == true {
				c.Send("Welcome")
				return nil
			} else {
				c.Send("Mật khẩu không chính xác")
				c.Send("Tùy chọn tiếp theo", helpers.Passcode_Menu_InlineKeys)
				return nil
			}
		}

		fmt.Println(user)
		return c.Send("Sorry, bot không xử lý các thông tin bạn gửi lên!")
	})

	b.Handle(tele.OnPhoto, func(c tele.Context) error {
		fmt.Println(tele.OnPhoto)
		user := c.Sender()
		fmt.Println(user)
		return c.Send("Sorry, bot không xử lý file ảnh bạn gửi lên!")
	})

	b.Handle(tele.OnPoll, func(c tele.Context) error {
		fmt.Println(tele.OnPoll)
		user := c.Sender()
		fmt.Println(user)
		return c.Send("Sorry, bot không xử lý bảng poll bạn gửi lên!")
	})
}
