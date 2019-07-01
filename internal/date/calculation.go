package date

import (
	"fmt"
	"time"
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

func GetFullDate(time time.Time)  string{
	return fmt.Sprintf("%s, %d %v %d",time.Weekday(),time.Day(),time.Month(),time.Year())
}

func (duration *Duration)GetDays(){
	duration.Days = duration.Hours/24
}

func (duration *Duration)GetWeeks()  {
	duration.Weeks.TotalWeeks = duration.Days/7
	duration.Weeks.DaysOfWeek = duration.Days%7
}

func GetMonths(startTime,endTime time.Time) Months {
	diffYear := endTime.Year()-startTime.Year()
	diffMonth := int(endTime.Month()-startTime.Month())+(diffYear*12)
	diffDay := endTime.Day()-startTime.Day()
	if diffDay<0 {
		diffMonth--
		diffDay+=31
	}
	return Months{
		TotalMonths:diffMonth,
		DaysOfMonth:diffDay,
	}
}