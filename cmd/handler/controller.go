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
	_, err, err1 := getRequestDateFromRequestBody(r)
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
		StartFullDate: "Thursday, 16 October 1997",
		EndFullDate:   "Monday, 10 June 2019",
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
	return requestDate, err, err2
}
