package date

import (
	"fmt"
	"time"
)

const (
	hoursOfDay   = 24
	daysOfWeek   = 7
	monthsOfYear = 12
	daysOfMonth  = 31
)

type Weeks struct {
	TotalWeeks int `json:"total_weeks"`
	DaysOfWeek int `json:"days_of_week"`
}

type Months struct {
	TotalMonths int `json:"total_months"`
	DaysOfMonth int `json:"days_of_month"`
}

type Duration struct {
	Seconds       int    `json:"seconds"`
	Minutes       int    `json:"minutes"`
	Hours         int    `json:"hours"`
	Days          int    `json:"days"`
	StartFullDate string `json:"start_full_date"`
	EndFullDate   string `json:"end_full_date"`
	Weeks         Weeks  `json:"weeks"`
	Months        Months `json:"months"`
}

func GetFullDate(time time.Time) string {
	return fmt.Sprintf("%s, %d %v %d", time.Weekday(), time.Day(), time.Month(), time.Year())
}

func (duration *Duration) getDays() {
	duration.Days = duration.Hours / hoursOfDay
}

func (duration *Duration) getWeeks() {
	duration.Weeks.TotalWeeks = duration.Days / daysOfWeek
	duration.Weeks.DaysOfWeek = duration.Days % daysOfWeek
}

func getMonths(startTime, endTime time.Time) Months {
	diffYear := endTime.Year() - startTime.Year()
	diffMonth := int(endTime.Month()-startTime.Month()) + (diffYear * monthsOfYear)
	diffDay := endTime.Day() - startTime.Day()
	if diffDay < 0 {
		diffMonth--
		diffDay += daysOfMonth
	}
	return Months{
		TotalMonths: diffMonth,
		DaysOfMonth: diffDay,
	}
}

func CalculateDuration(startTime, endTime time.Time) Duration {
	diffTime := endTime.Sub(startTime)
	duration := Duration{
		Seconds: int(diffTime.Seconds()),
		Minutes: int(diffTime.Minutes()),
		Hours:   int(diffTime.Hours()),
	}
	duration.getDays()
	duration.getWeeks()
	duration.Months = getMonths(startTime, endTime)
	return duration
}
