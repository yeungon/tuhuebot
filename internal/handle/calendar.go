package handle

import (
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

var calendar_url string = "https://www.dhsphue.edu.vn/lichxem.aspx"

func FetchCalendar() {
	// Launch a new browser with custom flags using rod.Launcher
	url := launcher.New().NoSandbox(true).MustLaunch()
	// Connect to the launched browser using ControlURL
	browser := rod.New().ControlURL(url).MustConnect()
	page := browser.MustPage(calendar_url).MustWindowFullscreen()

	page.MustWaitStable().MustScreenshot("vanban/" + FileNameCalendar())

}

func FileNameCalendar() string {
	timeLoc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	currentTime := time.Now().In(timeLoc)
	file_Name := fmt.Sprintf("%v_%v_%v.png", currentTime.Year(), currentTime.Month(), currentTime.Weekday())
	return file_Name
}

func Calendar(b *tele.Bot) {
	b.Handle(&helpers.Calendar_Tracking, func(c tele.Context) error {
		if helpers.IsAdmin(c) == false {
			return nil
		}
		var fileName = "vanban/" + FileNameCalendar()
		//var fileName = "vanban/" + FileName()
		if _, err := os.Stat(fileName); err == nil {
			fmt.Println("File exists:", fileName)
			c.Send("Đây là lịch công tác theo tuần, ảnh chụp tại thời điểm gần đây nhất. Thời gian: " + TimeFetch())
		} else if os.IsNotExist(err) {
			c.Send("Đang truy cập và lấy dữ liệu văn bản. Xin đợi... Thời gian: " + TimeFetch())
			FetchCalendar()
			fmt.Println("File does not exist:", fileName)
		} else {
			fmt.Println("Error checking file:", err)
		}

		a := &tele.Photo{File: tele.FromDisk(fileName)}

		fmt.Println(FileNameCalendar())
		c.Send(a)

		return nil
	})
}
