package date

import (
	"fmt"
	"time"
)

type YearMonthDay struct {
	year  int `json:"year"`
	month int `json:"month"`
	day   int `json:"day"`
}

type Duration struct {
	seconds int `json:"seconds"`
	minutes int `json:"minutes"`
	hours int `json:"hours"`
	days int `json:"days"`
	weeks WeeksDays `json:"weeks"`
	months MonthsDays `json:"months"`
}
type WeeksDays struct {
	weeks int `json:"weeks"`
	days int `json:"days"`
}

type MonthsDays struct {
	months int `json:"months"`
	days int `json:"days"`
}

func (date YearMonthDay) GetFullDate() string {
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

func CalculateMonths(startDate YearMonthDay, endDate YearMonthDay) MonthsDays {
	diffYear := endDate.year - startDate.year
	diffMonth := endDate.month - startDate.month + (diffYear * 12)
	diffDay := endDate.day - startDate.day
	if diffDay<0 {
		diffDay+=31
		diffMonth--
	}
	return MonthsDays{
		months: diffMonth,
		days:   diffDay,
	}
}
