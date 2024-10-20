package handle

import (
	"fmt"
	"strings"

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

func getName(c tele.Context) string {
	name := c.Sender().FirstName
	if len(name) > 0 {
		return name
	}
	return c.Sender().Username
}

func Passcode(b *tele.Bot) {
	b.Handle(tele.OnText, func(c tele.Context) error {
		user := c.Sender().ID
		db := sqlite.DB()
		current_user := users.GetCurrentUser(db, user)

		if current_user.State == true {
			user_input := strings.TrimSpace(c.Text())
			passcode := check_passcode(user_input)
			users.SetUserState(db, user, false)

			if passcode == true {
				users.SetUserLevel(db, user, 2)
				c.Send("Welcome " + getName(c) + ". Đây là khu vực truy cập dữ liệu level 2, hỗ trợ đa dạng hơn các nghiệp vụ xử lý thông tin. Bạn sẽ không cần nhập mật khẩu khi truy xuất.")
				return nil
			} else {
				c.Send("😮‍💨 Mật khẩu không chính xác!")
				c.Send("Tùy chọn tiếp theo 👇", helpers.Passcode_Menu_InlineKeys)
				return nil
			}
		}

		if current_user.StateAsking == true {
			user_input := strings.TrimSpace(c.Text())
			users.SetUserStateAsking(db, user, false)
			c.Send("Bot đã nhận được câu hỏi của bạn! Bot sẽ cập nhật dữ liệu khi có được câu trả lời của thầy cô, bộ phận chuyên môn!\n")
			//TODO: Save to xata
			fmt.Println(user_input)
			return nil
		}

		fmt.Println(user)
		return c.Send("Sorry, bot không xử lý các thông tin bạn gửi lên!")
	})

	b.Handle(&helpers.Back_To_Main_Menu, func(c tele.Context) error {
		intro := "🅰🅱©↩📧🎏⛽♓ℹ🗾🎋👢Ⓜ♑⭕🅿♌⚡🌴⛎✌Ⓩ"
		return c.Send(intro, helpers.MainMenu_InlineKeys)
	})

	b.Handle(&helpers.Keep_Typing_Passcode, func(c tele.Context) error {
		intro := "Xin nhập lại mật khẩu!"
		user := c.Sender().ID
		db := sqlite.DB()
		users.SetUserState(db, user, true)
		return c.Send(intro)
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
