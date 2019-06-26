package date

import (
	"fmt"
	"time"
)

type YearMonthDay struct {
	Year  int `json:"Year"`
	Month int `json:"Month"`
	Day   int `json:"Day"`
}

type Duration struct {
	Seconds int        `json:"Seconds"`
	Minutes int        `json:"Minutes"`
	Hours   int        `json:"Hours"`
	Days    int        `json:"Days"`
	Weeks   WeeksDays  `json:"Weeks"`
	Months  MonthsDays `json:"Months"`
}
type WeeksDays struct {
	Weeks int `json:"Weeks"`
	Days  int `json:"Days"`
}

type MonthsDays struct {
	Months int `json:"Months"`
	Days   int `json:"Days"`
}

func (date YearMonthDay) GetFullDate() string {
	dateTime := time.Date(date.Year, time.Month(date.Month), date.Day, 0, 0, 0, 0, time.UTC)
	return fmt.Sprintf("%s, %v %s %v", dateTime.Weekday(), dateTime.Day(), dateTime.Month(), dateTime.Year())
}

func CalculateDurationStartTimeToEndTime(startDate YearMonthDay, endDate YearMonthDay) Duration {
	start := time.Date(startDate.Year, time.Month(startDate.Month), startDate.Day, 0, 0, 0, 0, time.UTC)
	end := time.Date(endDate.Year, time.Month(endDate.Month), endDate.Day, 0, 0, 0, 0, time.UTC)
	diff := end.Sub(start)
	return Duration{
		Seconds: int(diff.Seconds()),
		Minutes: int(diff.Minutes()),
		Hours:   int(diff.Hours()),
	}
}

func (d Duration) GetDays() int {
	return d.Hours / 24
}

func (d Duration) GetWeeks() WeeksDays {
	return WeeksDays{
		Weeks: d.Days / 7,
		Days:  d.Days % 7,
	}
}

func CalculateMonths(startDate YearMonthDay, endDate YearMonthDay) MonthsDays {
	diffYear := endDate.Year - startDate.Year
	diffMonth := endDate.Month - startDate.Month + (diffYear * 12)
	diffDay := endDate.Day - startDate.Day
	if diffDay<0 {
		diffDay+=31
		diffMonth--
	}
	return MonthsDays{
		Months: diffMonth,
		Days:   diffDay,
	}
}
