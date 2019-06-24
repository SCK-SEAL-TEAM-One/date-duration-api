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

type Duration struct{
	seconds int
	minutes int
	hours int
	days int
	weeks WeeksDays
	months MonthsDays
}

type MonthsDays struct{
	months int
	days int
}

type WeeksDays struct{
	weeks int
	days int
}

func (date YearMonthDay) GetFullDate() string {
	dateTime := time.Date(date.year,time.Month(date.month),date.day,0,0,0,0,time.UTC)
	return fmt.Sprintf("%s, %d %s %d",dateTime.Weekday(),date.day,dateTime.Month(),date.year)
}

func CalculateDurationStartTimeToEndTime(start YearMonthDay,end YearMonthDay) Duration{
	startTime := time.Date(start.year,time.Month(start.month),start.day,0,0,0,0,time.UTC)
	endTime := time.Date(end.year,time.Month(end.month),end.day,0,0,0,0,time.UTC)
	duration := endTime.Sub(startTime) 
	return Duration{
		seconds:int(duration.Seconds()),
		minutes:int(duration.Minutes()),
		hours:int(duration.Hours()),
	}
}