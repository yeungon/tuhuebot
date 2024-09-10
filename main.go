package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Hi, your bot is running as expected!")

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hi, chào mừng bạn đến với bot trợ lý học tập KTU. Chúc bạn một ngày tốt lành. Xin lựa chọn tác vụ phía dưới")
	})

	b.Handle("/about", func(c tele.Context) error {
		return c.Send("Đây là phần mềm dạng bot hỗ trợ tự động, giúp sinh viên tiếp cận thông tin học tập nhanh và thuận lợi hơn. Thông tin mang giá trị tham khảo. Hãy liên hệ trực tiếp với giảng viên hướng dẫn hoặc cố vấn học tập để có được trợ giúp chi tiết hơn!")
	})

	b.Handle("/qa", func(c tele.Context) error {
		return c.Send("Hỏi đáp")
	})

	b.Handle("/guide", func(c tele.Context) error {
		return c.Send("Hướng dẫn")
	})

	b.Handle("/help", func(c tele.Context) error {
		return c.Send("Các gợi ý giúp bạn sử dụng bot tốt hơn")
	})

	b.Start()
}
