package admin

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"math/rand"
	"strconv"
	"time"

	"github.com/yeungon/tuhuebot/internal/database/bbolt"
	"github.com/yeungon/tuhuebot/internal/database/pg"
	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/users"
	"github.com/yeungon/tuhuebot/pkg/cache"
	"github.com/yeungon/tuhuebot/pkg/helpers"
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

	b.Handle("/user", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		user := c.Sender().ID
		db := sqlite.DB()
		current_user := users.GetCurrentUser(db, user)

		users.SetUserState(db, user, false)

		state := users.UserState(db, user)

		info := fmt.Sprintf("%v - %s - trạng thái mở hay đóng %v", *current_user.Username, current_user.FirstName, state)
		c.Send(info)
		c.Send("Hello, testing user")
		return nil
	})

	b.Handle("/get", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		//user := c.Sender()
		// helpers.PrintStruct(user)
		pgdata := pg.PG()
		// events := pg.GetEvent(pgdata)

		newQA := &pg.QA{
			UserAsked:     "1",         // Example user ID
			UserAnswered:  "some_user", // Example answerer
			Question:      "hello What is the capital of France?",
			Published:     false, // Default to false
			XataCreatedat: time.Now(),
			XataUpdatedat: time.Now(),
		}

		pg.CreateQA(pgdata, newQA)

		// db := sqlite.DB()
		// usersList := users.GetAllUser(db)
		// first := usersList[0]

		// fmt.Println(first)

		// // Print the retrieved users.
		// for _, event := range events {
		// 	event := fmt.Sprintf("%d - %s ", event.Month, event.EventData)
		// 	c.Send(event)
		// }

		fmt.Println("Testing bun ORM - Get data from sql")
		return nil
	})

	b.Handle("/log", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}

		// Log some messages with structured data
		slog.Debug("This is a debug message", "module", "auth", "status", "success")
		slog.Info("User logged in", "user", "john_doe", "module", "auth")
		slog.Warn("Disk space is running low", "module", "storage", "free_space", "5GB")
		slog.Error("Failed to connect to database", "module", "db", "error", "connection timeout")

		fmt.Println("Testing log")
		return nil
	})
	b.Handle("/bbolt", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		fmt.Println("testing kv bbolt")

		random := rand.Intn(100)

		byteData := []byte(strconv.Itoa(random))

		bbolt.UserAdd([]byte("12345"), byteData)
		return nil
	})

	b.Handle("/read", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		bbolt.PrintAllKeyValues("users")
		return nil
	})

	b.Handle("/add", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		current_user := c.Sender()
		current_user_kv := fmt.Sprintf("%v", current_user.ID)
		bbolt.UserAdd([]byte(current_user_kv), IntToBytes(100))

		return nil
	})

	b.Handle("/clean", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		bbolt.CleanAllKeyValues("users")
		fmt.Println("Clean up the k-v")
		return nil
	})

	b.Handle("/admin", func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		return nil
	})
}

func IntToBytes(n int) []byte {
	b := make([]byte, 8) // Use 8 bytes for int64 or uint64
	binary.BigEndian.PutUint64(b, uint64(n))
	return b
}
