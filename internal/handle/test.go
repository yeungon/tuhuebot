package handle

import (
	"fmt"
	"math/rand"
	"strconv"

	tele "gopkg.in/telebot.v3"
)

func Test(b *tele.Bot) {
	random := rand.Intn(100)

	fmt.Println(random)

	callback_random := strconv.Itoa(random)
	// Create an inline button with data
	inlineBtn := tele.InlineButton{
		Unique: "my_callback",
		Text:   "Click me",
		Data:   callback_random, // Set the callback data
	}

	// Handle the callback for the button
	b.Handle(&inlineBtn, func(c tele.Context) error {
		random := rand.Intn(100)
		callback_random := strconv.Itoa(random)

		// Create a new button with updated random number as data <-- update the previous inline button instead of creating a new one
		inlineBtn = tele.InlineButton{
			Unique: "my_callback",
			Text:   "Click me again!",
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
		// Create a reply with the inline button
		inlineKeys := [][]tele.InlineButton{
			{inlineBtn},
		}
		return c.Send("Click the button below:", &tele.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		})
	})

	// b.Handle("/qa", func(c tele.Context) error {
	// 	return c.Reply("hỏi đáp")
	// })

	// b.Handle(&helpers.QA, func(c tele.Context) error {
	// 	return c.Send("hỏi đáp")
	// })
}
