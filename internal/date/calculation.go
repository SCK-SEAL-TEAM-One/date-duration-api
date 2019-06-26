package date

import (
	"fmt"
	"time"
)

const HoursInDay = 24
const DaysInWeek = 7
const DaysInMonth = 31
const MonthsInYear = 12

type YearMonthDay struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type Duration struct {
	Seconds int        `json:"seconds"`
	Minutes int        `json:"minutes"`
	Hours   int        `json:"hours"`
	Days    int        `json:"days"`
	Weeks   WeeksDays  `json:"weeks"`
	Months  MonthsDays `json:"months"`
}
type WeeksDays struct {
	Weeks int `json:"weeks"`
	Days  int `json:"days"`
}

type MonthsDays struct {
	Months int `json:"months"`
	Days   int `json:"days"`
}

func (date YearMonthDay) GetFullDate() string {
	dateTime := time.Date(date.Year, time.Month(date.Month), date.Day, 0, 0, 0, 0, time.UTC)
	return fmt.Sprintf("%s, %v %s %v", dateTime.Weekday(), dateTime.Day(), dateTime.Month(), dateTime.Year())
}

func CalculateDurationStartTimeToEndTime(startDate YearMonthDay, endDate YearMonthDay) (duration Duration) {
	start := time.Date(startDate.Year, time.Month(startDate.Month), startDate.Day, 0, 0, 0, 0, time.UTC)
	end := time.Date(endDate.Year, time.Month(endDate.Month), endDate.Day, 0, 0, 0, 0, time.UTC)
	diff := end.Sub(start)
	duration = Duration{
		Seconds: int(diff.Seconds()),
		Minutes: int(diff.Minutes()),
		Hours:   int(diff.Hours()),
	}
	duration.Days = duration.GetDays()
	duration.Weeks = duration.GetWeeks()
	duration.Months = CalculateMonths(startDate, endDate)
	return duration
}

func (d Duration) GetDays() int {
	return d.Hours / HoursInDay
}

func (d Duration) GetWeeks() WeeksDays {
	return WeeksDays{
		Weeks: d.Days / DaysInWeek,
		Days:  d.Days % DaysInWeek,
	}
}

func CalculateMonths(startDate YearMonthDay, endDate YearMonthDay) MonthsDays {
	diffYear := endDate.Year - startDate.Year
	diffMonth := endDate.Month - startDate.Month + (diffYear * MonthsInYear)
	diffDay := endDate.Day - startDate.Day
	if diffDay < 0 {
		diffDay += DaysInMonth
		diffMonth--
	}
	return MonthsDays{
		Months: diffMonth,
		Days:   diffDay,
	}
}
