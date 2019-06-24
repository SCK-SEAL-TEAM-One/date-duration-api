package date

import (
	"time"
	"fmt"
)

type YearMonthDay struct {
	year int
	month int
	day int
}



func (date YearMonthDay) GetFullDate() string {
	dateTime := time.Date(date.year,time.Month(date.month),date.day,0,0,0,0,time.UTC)
	return fmt.Sprintf("%s, %d %s %d",dateTime.Weekday(),date.day,dateTime.Month(),date.year)
}