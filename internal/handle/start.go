package handle

import (
	"fmt"

	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/users"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func Name(firstName, username string) string {
	if len(firstName) == 0 {
		return username
	} else {
		return firstName
	}
}

func Start(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		var (
			user    = c.Sender()
			intro   = "Welcome onboard"
			welcome = "Chào mừng bạn đến với bot hỗ trợ học tập tự động. Chúc bạn một ngày tốt lành.\n\nDưới đây là các chức năng, thông tin chính hiện có của bot:"
		)

		// Extract user details
		firstName := user.FirstName
		lastName := user.LastName
		username := user.Username
		telegramUserID := user.ID
		isBot := user.IsBot

		var lastNamePtr *string
		if lastName != "" {
			lastNamePtr = &lastName // Set lastName if it's not empty
		}

		name := Name(firstName, username)
		introduction := fmt.Sprintf("%s %s. %s", intro, name, welcome)
		// Store the user in the database, TODO: check the user if it has been stored.

		db := sqlite.DB()
		users_data := []*users.User{
			{TelegramUserID: telegramUserID,
				FirstName: firstName,
				LastName:  lastNamePtr,
				Username:  &username, // Username is also a pointer
				IsBot:     isBot,
			},
		}

		// Store the user in the database (consider checking if the user already exists)
		users.CreateUser(db, users_data)

		c.Send(introduction, helpers.MainMenu_InlineKeys)
		return nil
	})
}
