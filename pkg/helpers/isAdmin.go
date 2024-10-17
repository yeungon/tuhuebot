package helpers

import (
	"log"
	"strconv"

	"github.com/yeungon/tuhuebot/internal/config"
	tele "gopkg.in/telebot.v3"
)

func IsAdmin(c tele.Context) bool {
	admin_id := config.Get().AdminID
	current_user := c.Sender()
	current_user_id := strconv.Itoa(int(current_user.ID))
	if admin_id != current_user_id {
		log.Println("Current user is not the admin")
		return false
	}

	log.Println("admin is using the feature")
	return true
}
