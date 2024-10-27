package handle

import (
	"fmt"
	"strconv"

	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/users"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func HandleAssistant(c tele.Context) error {
	user := c.Sender().ID
	db := sqlite.DB()
	current_user := users.GetCurrentUser(db, user)

	userInfo := current_user.FirstName
	if len(userInfo) == 0 {
		userInfo = *current_user.Username
	}
	if len(userInfo) == 0 {
		userInfo = strconv.FormatInt(current_user.ID, 10)
	}

	if current_user.Level > 1 {
		message := fmt.Sprintf("Xin chào %v, các tùy chọn nghiệp vụ:", userInfo)
		c.Send(message, helpers.Assitant_InlineKeys)
		return nil
	}

	content := "Đây là góc hỗ trợ trợ lý. Bạn cần có mật khẩu để truy xuất dữ liệu! Bạn chỉ cần nhập khẩu đúng một lần duy nhất.\n\nXin nhập mật khẩu: "
	users.SetUserState(db, user, true)
	c.Send(content)
	return nil
}

func Assistant(b *tele.Bot) {
	b.Handle("/assistant", func(c tele.Context) error {
		HandleAssistant(c)
		return nil
	})

	b.Handle(&helpers.Assistant, func(c tele.Context) error {
		HandleAssistant(c)
		return nil
	})

}
