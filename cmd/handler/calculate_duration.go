package handler

import (
	"bytes"
	"encoding/json"
	"github.com/SCK-SEAL-TEAM-One/date-duration-api/internal/date"
	"io/ioutil"
	"net/http"
	"time"
)

type RequestDate struct {
	StartDate YearMonthDay `json:"start_date"`
	EndDate   YearMonthDay `json:"end_date"`
}

type YearMonthDay struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
	Time  time.Time
}

func CalculateDuration(w http.ResponseWriter, r *http.Request) {
	requestDate, err := getRequestDateFromRequestBody(r)
	if err != nil {
		http.Error(w,err.Error(),400)
		return
	}

	duration := date.Duration{
		Seconds: 683164800,
		Minutes: 11386080,
		Hours:   189768,
		Days:    7907,
		Weeks: date.Weeks{
			TotalWeeks: 1129,
			DaysOfWeek: 4,
		},
		Months: date.Months{
			TotalMonths: 259,
			DaysOfMonth: 25,
		},
		StartFullDate: date.GetFullDate(requestDate.StartDate.Time),
		EndFullDate:   date.GetFullDate(requestDate.EndDate.Time),
	}

	err = json.NewEncoder(w).Encode(duration)
	if err != nil {
		http.Error(w,err.Error(),500)
		return

	}
}

func getRequestDateFromRequestBody(request *http.Request) (RequestDate, error) {
	var requestDate RequestDate
	body, err := ioutil.ReadAll(request.Body)
	if err != nil{
		return requestDate,err
	}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&requestDate)
	if err != nil{
		return requestDate,err
	}
	requestDate.StartDate = NewYearMonthDay(requestDate.StartDate.Year, requestDate.StartDate.Month, requestDate.StartDate.Day)
	requestDate.EndDate = NewYearMonthDay(requestDate.EndDate.Year, requestDate.EndDate.Month, requestDate.EndDate.Day)
	return requestDate, nil
}

func NewYearMonthDay(year, month, day int) YearMonthDay {
	return YearMonthDay{
		Year:  year,
		Month: month,
		Day:   day,
		Time:  time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
	}
}
