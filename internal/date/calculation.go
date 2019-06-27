package date

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
