package handle

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/yeungon/tuhuebot/internal/database/bbolt"
	"github.com/yeungon/tuhuebot/internal/database/pg"
	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/users"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func IntToBytes(n int) []byte {
	b := make([]byte, 8) // Use 8 bytes for int64 or uint64
	binary.BigEndian.PutUint64(b, uint64(n))
	return b
}

func BytesToInt(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func UpdateOffset(c tele.Context, stepstate int, possible_offset int) int {
	current_user := c.Sender()
	current_user_kv := fmt.Sprintf("%v", current_user.ID)
	current_user_pagination := bbolt.UserRead([]byte(current_user_kv))

	var currentState int

	if len(current_user_pagination) == 0 {
		// If no value exists, initialize it to 0.
		fmt.Println("current_user_pagination empty byte[]")
		currentState = 0
	} else {
		// Convert the existing byte array to an integer.
		currentState = BytesToInt(current_user_pagination)
		fmt.Printf("Current state as int: %d\n", currentState)
	}

	// Calculate the new state by adding the step.
	newState := currentState + stepstate
	if newState < 0 {
		newState = 0
	}

	if newState >= possible_offset {
		newState = possible_offset
	}

	// Defer the update so that it happens after returning the new state.
	defer func() {
		fmt.Printf("Saving new state: %d\n", newState)
		bbolt.UserAdd([]byte(current_user_kv), IntToBytes(newState))
	}()

	// Return the new state.
	return newState
}

func FetchQAPG(b *tele.Bot, c tele.Context, control int) {
	current_user := c.Sender()

	pgdata := pg.PG()
	question_answer := pg.GetQuestionAnswer(pgdata)
	total_qa := len(question_answer)
	possible_offset := total_qa / 5

	inform_Message := fmt.Sprintf("Các câu hỏi thường gặp. Hiện có %d câu hỏi.", total_qa)

	c.Send(inform_Message)

	fmt.Printf("Tổng số qa:  %d\n", total_qa)
	fmt.Printf("possible_offset:  %d\n", possible_offset)

	offset := UpdateOffset(c, control, possible_offset)

	fmt.Printf("Offset hiện tại: %d \n", offset)

	step := 5 * offset
	starting := 0 + step
	ending := starting + 5

	if ending >= total_qa {
		ending = total_qa
		starting = ending - 5
		step = ending
	}

	if starting < 0 {
		starting = 0
		ending = starting + 5
		step = 0
	}

	portion_slice := question_answer[starting:ending]

	fmt.Println("starting-ending: ", starting, ending)

	for index, record := range portion_slice {
		if record.Published == true {
			index_string := strconv.Itoa(index + starting + 1)
			questionMsg := "🌓 🅀🅄🄴🅂🅃🄸🄾🄽 <i>" + index_string + ": " + record.Question + "</i>"
			b.Send(current_user, questionMsg, &tele.SendOptions{
				ParseMode: "HTML",
			})
			// answerMsgTexta := "<b>✅ 🄰🄽🅂🅆🄴🅁: </b>" + record.Answer
			answerMsgTexta := "<b>✅ </b>" + *record.Answer
			b.Send(current_user, answerMsgTexta, &tele.SendOptions{
				ParseMode: "HTML",
			})
		}
	}
}

func Qa(b *tele.Bot) {
	b.Handle("/qa", func(c tele.Context) error {
		FetchQAPG(b, c, 0)
		c.Send("Xem các câu hỏi khác", helpers.QA_Menu_InlineKeys)
		return nil
	})

	b.Handle(&helpers.QA, func(c tele.Context) error {
		FetchQAPG(b, c, 0)
		c.Send("Xem các câu hỏi khác", helpers.QA_Menu_InlineKeys)
		return nil
	})
	ControlQuestion(b)
	PostQuestion(b)
}

func ControlQuestion(b *tele.Bot) {
	b.Handle(&helpers.Back_QA, func(c tele.Context) error {
		fmt.Println("Control -")
		FetchQAPG(b, c, -1)
		c.Send("Xem các câu hỏi khác", helpers.QA_Menu_InlineKeys)
		return nil
	})

	b.Handle(&helpers.Forward_QA, func(c tele.Context) error {
		fmt.Println("Control +")
		FetchQAPG(b, c, 1)
		c.Send("Xem các câu hỏi khác", helpers.QA_Menu_InlineKeys)
		return nil
	})
}

func PostQuestion(b *tele.Bot) {
	b.Handle(&helpers.Post_QA, func(c tele.Context) error {
		// c.Send(current)
		user_id := c.Sender().ID
		user := c.Sender().ID
		db := sqlite.DB()
		users.SetUserStateAsking(db, user_id, true)
		current_user_asking := users.UserStateAsking(db, user)
		fmt.Println(current_user_asking)
		c.Send("Bot đã bật chế độ nhận câu hỏi. Xin đặt câu hỏi! 🔓")
		return nil
	})
}
