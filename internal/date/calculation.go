package date

import (
	"fmt"
	"time"
)

const (
	hoursOfDay   = 24
	daysOfWeek   = 7
	daysOfMonth  = 31
	monthsOfYear = 12
)

type Duration struct {
	Seconds       int    `json:"seconds"`
	Minutes       int    `json:"minutes"`
	Hours         int    `json:"hours"`
	Days          int    `json:"days"`
	Weeks         Weeks  `json:"weeks"`
	Months        Months `json:"months"`
	StartFullDate string `json:"start_full_date"`
	EndFullDate   string `json:"end_full_date"`
}

type Weeks struct {
	TotalWeeks int `json:"total_weeks"`
	DaysOfWeek int `json:"days_of_week"`
}

type Months struct {
	TotalMonths int `json:"total_months"`
	DaysOfMonth int `json:"days_of_month"`
}

func DurationBetween(startTime, endTime time.Time) Duration {
	duration := endTime.Sub(startTime)
	days := getDays(int(duration.Hours()))
	return Duration{
		Seconds:       int(duration.Seconds()),
		Minutes:       int(duration.Minutes()),
		Hours:         int(duration.Hours()),
		Days:          days,
		Weeks:         getWeeks(days),
		Months:        getMonths(startTime, endTime),
		StartFullDate: getFullDate(startTime),
		EndFullDate:   getFullDate(endTime),
	}
}

func getDays(hour int) int {
	days := hour / hoursOfDay
	return days
}

func getWeeks(day int) Weeks {
	return Weeks{
		TotalWeeks: day / daysOfWeek,
		DaysOfWeek: day % daysOfWeek,
	}
}

func getMonths(startTime, endTime time.Time) Months {
	diffYears := endTime.Year() - startTime.Year()
	diffMonths := int(endTime.Month()-startTime.Month()) + (diffYears * monthsOfYear)
	diffDays := endTime.Day() - startTime.Day()
	if diffDays < 0 {
		diffDays += daysOfMonth
		diffMonths--
	}
	return Months{TotalMonths: diffMonths, DaysOfMonth: diffDays}
}

func getFullDate(date time.Time) string {
	return fmt.Sprintf("%s, %d %s %d", date.Weekday(), date.Day(), time.Month(date.Month()), date.Year())
}
