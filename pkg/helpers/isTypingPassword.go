package helpers

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

func IsTypingPassword(c tele.Context) {
	current_user := c.Sender()
	fmt.Println(current_user)
}
