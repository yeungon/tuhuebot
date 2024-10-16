package handle

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/users"
	"github.com/yeungon/tuhuebot/pkg/cache"
	"github.com/yeungon/tuhuebot/pkg/helpers"
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
		if helpers.IsAdmin(c) == false {
			return nil
		}

		// test cache
		cache.TestCache()

		// Create a reply with the inline button
		inlineKeys := [][]tele.InlineButton{
			{inlineBtn},
		}
		return c.Send("Click the button below:", &tele.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		})
	})

	b.Handle("/state", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		//fmt.Println(c.Message().Text)
		hello := "Hello world"
		reference.Test()

		answerMsgText := "<b>Đang test state management</b>🍟" + hello
		b.Send(c.Sender(), answerMsgText, &tele.SendOptions{
			ParseMode: "HTML",
		})
		return nil
	})

	b.Handle("/bun", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		db := sqlite.DB()
		users_data := []*users.User{
			{FirstName: "Alice", TelegramUserID: 30},
		}

		users.CreateUser(db, users_data)

		fmt.Println("Testing bun ORM")
		return nil
	})

	b.Handle("/get", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		user := c.Sender()
		helpers.PrintStruct(user)

		db := sqlite.DB()
		usersList := users.GetAllUser(db)
		first := usersList[0]

		fmt.Println(first)

		// Print the retrieved users.
		for _, user := range usersList {
			users := fmt.Sprintf("%d - %s - %d", user.ID, user.FirstName, user.TelegramUserID)
			c.Send(users)
		}

		fmt.Println("Testing bun ORM - Get data from sql")
		return nil
	})

}
