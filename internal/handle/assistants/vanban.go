package assistants

import (
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/yeungon/tuhuebot/internal/config"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func FolderExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

func createFolder() {
	dirName := "vanban"
	exists, err := FolderExists(dirName)
	if err != nil {
		fmt.Println("Error:", err)
	} else if exists {
		fmt.Printf("Folder %s exists\n!", dirName)
	} else {
		// Create the directory
		err = os.Mkdir(dirName, 0755) // 0755 is the permission
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		fmt.Println("Directory created successfully:", dirName)
	}
}

func FileName() string {
	timeLoc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	currentTime := time.Now().In(timeLoc)
	file_Name := fmt.Sprintf("%v_%v_%v_%v.png", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour())
	return file_Name
}

func TimeFetch() string {
	timeLoc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	currentTime := time.Now().In(timeLoc)
	time_fetched := fmt.Sprintf("%v giờ, ngày %v, tháng %v, năm %v", currentTime.Hour(), currentTime.Day(), currentTime.Month(), currentTime.Year())
	return time_fetched
}

func HandleVanban(c tele.Context) {
	// Launch a new browser with custom flags using rod.Launcher
	url := launcher.New().NoSandbox(true).MustLaunch()
	// Connect to the launched browser using ControlURL
	browser := rod.New().ControlURL(url).MustConnect()
	page := browser.MustPage("https://qlvb.dhsphue.edu.vn").MustWindowFullscreen()

	var sph_username = config.Get().SPH_USERNAME
	var sph_password = config.Get().SPH_PASSWORD
	page.MustElement("#txt_uid").MustInput(sph_username)
	page.MustElement("#txt_pws").MustInput(sph_password)
	page.MustElement("#bt_login").MustClick()
	time.Sleep(300 * time.Microsecond)
	page.MustWaitStable().MustElement("#bt_main").MustClick()

	createFolder()

	page.MustWaitStable().MustScreenshot("vanban/" + FileName())
	// time.Sleep(time.Hour)

}

func Vanban(b *tele.Bot) {
	b.Handle(&helpers.Tracking_Announcement, func(c tele.Context) error {
		var fileName = "vanban/" + FileName()
		if _, err := os.Stat(fileName); err == nil {
			fmt.Println("File exists:", fileName)
			c.Send("Đây là file ảnh chụp tại thời điểm gần đây nhất. Thời gian: " + TimeFetch())
		} else if os.IsNotExist(err) {
			c.Send("Đang truy cập và lấy dữ liệu văn bản. Xin đợi... Thời gian: " + TimeFetch())
			HandleVanban(c)
			fmt.Println("File does not exist:", fileName)
		} else {
			fmt.Println("Error checking file:", err)
		}
		a := &tele.Photo{File: tele.FromDisk(fileName)}
		c.Send(a)
		return nil
	})
}
