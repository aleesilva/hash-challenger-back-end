package helpers

import (
	"fmt"
	"time"
)

func VerifyBlackFridayDate() bool {
	time := time.Now()
	formattedDate := fmt.Sprintf("%02d/%02d", time.Month(), time.Day())
	bfDate := "12/31"
	if formattedDate != bfDate {
		return false
	} else {
		return true
	}
}
