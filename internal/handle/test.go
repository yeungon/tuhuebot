package handle

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/yeungon/tuhuebot/internal/config"
	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	"github.com/yeungon/tuhuebot/pkg/cache"
	"github.com/yeungon/tuhuebot/pkg/reference"
	tele "gopkg.in/telebot.v3"
)

func Test(b *tele.Bot) {
	random := rand.Intn(100)

	fmt.Println(random)

	callback_random := strconv.Itoa(random)
	// Create an inline button with data
	inlineBtn := tele.InlineButton{
		Unique: "my_callback",
		Text:   callback_random,
		Data:   callback_random, // Set the callback data
	}

	// Handle the callback for the button
	b.Handle(&inlineBtn, func(c tele.Context) error {
		random := rand.Intn(100)
		callback_random := strconv.Itoa(random)

		// <-- update the previous inline button instead of creating a new one
		inlineBtn = tele.InlineButton{
			Unique: "my_callback",
			Text:   callback_random,
			Data:   callback_random,
		}

		// Respond to the user
		c.Send("You clicked the button! Random number: " + callback_random)

		// Send a new button with updated data
		inlineKeys := [][]tele.InlineButton{
			{inlineBtn},
		}

		return c.Send("Click the button below for a new random number:", &tele.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		})
	})

	// Handle start command
	b.Handle("/test", func(c tele.Context) error {
		// test cache
		cache.TestCache()

		admin_id := config.Get().AdminID
		current_user := c.Sender()
		current_user_id := strconv.Itoa(int(current_user.ID))
		// only handle if the current user is the admin
		if admin_id != current_user_id {
			return nil
		}
		// Create a reply with the inline button
		inlineKeys := [][]tele.InlineButton{
			{inlineBtn},
		}
		return c.Send("Click the button below:", &tele.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		})
	})

	b.Handle("/html", func(c tele.Context) error {
		//fmt.Println(c.Message().Text)
		hello := "Hello world"
		answerMsgText := "<b>Trả lời (tham khảo)!</b>🍟" + hello
		b.Send(c.Sender(), answerMsgText, &tele.SendOptions{
			ParseMode: "HTML",
		})
		return nil
	})

	b.Handle("/state", func(c tele.Context) error {
		//fmt.Println(c.Message().Text)
		hello := "Hello world"
		reference.Test()

		answerMsgText := "<b>Đang test state management</b>🍟" + hello
		b.Send(c.Sender(), answerMsgText, &tele.SendOptions{
			ParseMode: "HTML",
		})
		return nil
	})

	b.Handle("/database", func(c tele.Context) error {
		//fmt.Println(c.Message().Text)
		hello := "Hello world"
		sqlite.SQLite()

		answerMsgText := "<b>Đang test sqlite database</b>🍟" + hello
		b.Send(c.Sender(), answerMsgText, &tele.SendOptions{
			ParseMode: "HTML",
		})
		return nil
	})

}
