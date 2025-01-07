package assistants

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	"github.com/yeungon/tuhuebot/internal/config"
	"github.com/yeungon/tuhuebot/internal/database/sqlite"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/users"
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

func StudentCheck(b *tele.Bot) {
	KeepCheckingStudent(b)
	PostCheckStudent(b)
}

func PostCheckStudent(b *tele.Bot) {
	b.Handle(&helpers.Students_Check, func(c tele.Context) error {
		// c.Send(current)
		user_id := c.Sender().ID
		user := c.Sender().ID
		db := sqlite.DB()
		users.SetUserStateChecking(db, user_id, true)
		current_user_checking := users.UserStateChecking(db, user)
		fmt.Println(current_user_checking)
		c.Send("Bot đã bật chế độ xem thông tin sinh viên. Xin nhập mã sinh viên! 💡")
		return nil
	})
}

func KeepCheckingStudent(b *tele.Bot) {
	b.Handle(&helpers.Keep_Checking_Student, func(c tele.Context) error {
		// c.Send(current)
		user_id := c.Sender().ID
		db := sqlite.DB()
		users.SetUserStateChecking(db, user_id, true)
		c.Send("Chế độ xem thông tin sinh viên đã được bật lại. Xin nhập mã sinh viên! 🌿")
		return nil
	})
}

func StudentCheckFetch(c tele.Context, studentID string) error {
	baseurl := config.Get().SPH_URL_ENDPOINT
	key := generateHash()
	endpoint := fmt.Sprintf("%s?key=%s&id=%s", baseurl, key, studentID)
	data, err := fetchData(endpoint)

	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return c.Send("Không tìm thấy thông tin, xin kiểm tra mã sinh viên. Vui lòng thử lại sau!")
	}

	fmt.Println("Fetched data successfully:")

	// Format the student info
	studentInfo := fmt.Sprintf(
		"Thông tin sinh viên:\nMã số: %s\nHọ và tên: %s\nLớp: %s\nSố điện thoại: %s\n",
		data.Info.FMasv, data.Info.FHoten, data.Info.FLop, data.Info.FPhone,
	)

	// Send student info first
	err = c.Send(studentInfo)
	if err != nil {
		log.Printf("Error sending student info: %v", err)
		return c.Send("Không thể gửi thông tin sinh viên.")
	}

	// Format scores (split if too long)
	const maxMessageLength = 4000 // Telegram's limit for a single message
	scores := "Kết quả học tập:\n"
	messages := []string{}
	for _, diem := range data.Diem {
		entry := fmt.Sprintf(
			"Học kỳ: %s | Môn học: %s (%s)\nĐiểm: %s | Điểm QT: %s | Điểm thi: %s | ĐVHT: %s\n\n",
			diem.Hk, diem.Tennh, diem.Mamh, diem.Diem, diem.Diemqt, diem.Diemthi, diem.Dvht,
		)

		if len(scores)+len(entry) > maxMessageLength {
			messages = append(messages, scores)
			scores = "Kết quả học tập:\n" // Start a new message
		}

		scores += entry
	}

	// Append the final message if there's remaining content
	if scores != "Kết quả học tập:\n" {
		messages = append(messages, scores)
	}

	// Send scores in parts
	for _, message := range messages {
		err = c.Send(message)
		if err != nil {
			log.Printf("Error sending scores: %v", err)
			return c.Send("Không thể gửi toàn bộ kết quả học tập. Vui lòng thử lại sau.")
		}
	}

	return nil

}

type ResponseData struct {
	Info struct {
		FMasv  string `json:"f_masv"`
		FHoten string `json:"f_hoten"`
		FLop   string `json:"f_lop"`
		FPhone string `json:"f_phone"`
	} `json:"info"`
	Diem []struct {
		Hk      string `json:"HK"`
		Mamh    string `json:"MAMH"`
		Diem    string `json:"DIEM"`
		Diemqt  string `json:"DIEMQT"`
		Diemthi string `json:"DIEMTHI"`
		Diemtl  string `json:"DIEMTL"`
		Tennh   string `json:"tennh"`
		Dvht    string `json:"dvht"`
	} `json:"diem"`
}

func fetchData(url string) (*ResponseData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	// Check if the status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var data ResponseData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return &data, nil
}

func generateHash() string {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	time.Local = loc
	prefix := config.Get().SECRET_FIRST
	suffix := config.Get().SECRET_SECOND
	currentTime := time.Now().In(loc)
	dateString := currentTime.Format("212006") // dmyyyy format
	input := prefix + dateString + suffix
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
