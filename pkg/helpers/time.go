package helpers

import (
	"fmt"
	"time"
)

func GetCurrentMonth() int {
	fmt.Println("currentMonth is call")
	timeLoc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	currentTime := time.Now().In(timeLoc)
	current_Month := int(currentTime.Month())
	return current_Month
}
