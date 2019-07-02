package date

import "time"

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

func getWeeks(day int) Weeks {
	return Weeks{
		TotalWeeks: day / 7,
		DaysOfWeek: day % 7,
	}
}

func getMonths(startTime, endTime time.Time) Months {
	diffYears := endTime.Year() - startTime.Year()
	diffMonths := int(endTime.Month()-startTime.Month()) + (diffYears * 12)
	diffDays := endTime.Day() - startTime.Day()
	if diffDays < 0 {
		diffDays += 31
		diffMonths--
	}
	return Months{TotalMonths: diffMonths, DaysOfMonth: diffDays}
}
