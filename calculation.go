package date

import (
	"fmt"
	"time"
)

type YearMonthDay struct {
	year  int
	month int
	day   int
}

func GetFullDate(date YearMonthDay) string {
	dateTime := time.Date(date.year, time.Month(date.month), date.day, 0, 0, 0, 0, time.UTC)
	return fmt.Sprintf("%s, %v %s %v", dateTime.Weekday(), dateTime.Day(), dateTime.Month(), dateTime.Year())
}
