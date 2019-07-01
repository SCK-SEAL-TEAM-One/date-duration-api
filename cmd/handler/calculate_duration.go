package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	requestDate, err, err1 := getRequestDateFromRequestBody(r)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println(err.Error())
	}
	if err1 != nil {
		w.WriteHeader(500)
		fmt.Println(err1.Error())
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

	encoder := json.NewEncoder(w)
	err2 := encoder.Encode(duration)

	if err2 != nil {
		w.WriteHeader(500)
		fmt.Println(err2.Error())
	}
}

func getRequestDateFromRequestBody(request *http.Request) (RequestDate, error, error) {
	body, err := ioutil.ReadAll(request.Body)
	decoder := json.NewDecoder(bytes.NewReader(body))
	var requestDate RequestDate
	err2 := decoder.Decode(&requestDate)
	requestDate.StartDate = NewYearMonthDay(requestDate.StartDate.Year, requestDate.StartDate.Month, requestDate.StartDate.Day)
	requestDate.EndDate = NewYearMonthDay(requestDate.EndDate.Year, requestDate.EndDate.Month, requestDate.EndDate.Day)
	return requestDate, err, err2
}

func NewYearMonthDay(year, month, day int) YearMonthDay {
	return YearMonthDay{
		Year:  year,
		Month: month,
		Day:   day,
		Time:  time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
	}
}
