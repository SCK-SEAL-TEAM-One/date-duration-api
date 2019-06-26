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

type WeeksDays struct {
	weeks int
	days  int
}

type MonthsDays struct {
	months int
	days   int
}

type Duration struct {
	seconds int
	minutes int
	hours   int
	days    int
	weeks   WeeksDays
	months  MonthsDays
}

func GetFullDate(date YearMonthDay) string {
	dateTime := time.Date(date.year, time.Month(date.month), date.day, 0, 0, 0, 0, time.UTC)
	return fmt.Sprintf("%s, %v %s %v", dateTime.Weekday(), dateTime.Day(), dateTime.Month(), dateTime.Year())
}

func CalculateDurationStartTimeToEndTime(startDate YearMonthDay, endDate YearMonthDay) Duration {
	start := time.Date(startDate.year, time.Month(startDate.month), startDate.day, 0, 0, 0, 0, time.UTC)
	end := time.Date(endDate.year, time.Month(endDate.month), endDate.day, 0, 0, 0, 0, time.UTC)
	diff := end.Sub(start)
	return Duration{
		seconds: int(diff.Seconds()),
		minutes: int(diff.Minutes()),
		hours:   int(diff.Hours()),
	}
}

func (d Duration) GetDays() int {
	return d.hours / 24
}

func (d Duration) GetWeeks() WeeksDays {
	return WeeksDays{
		weeks: d.days / 7,
		days:  d.days % 7,
	}
}
