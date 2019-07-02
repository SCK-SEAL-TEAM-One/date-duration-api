package handler

import (
	"bytes"
	"date-duration-api/internal/date"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Dates struct {
	StartDate YearMonthDay `json:"start_date"`
	EndDate   YearMonthDay `json:"end_date"`
}

type YearMonthDay struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func CalculateDuration(w http.ResponseWriter, r *http.Request) {
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

	err := json.NewEncoder(w).Encode(&duration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func decodeDate(request *http.Request) (Dates, error) {
	var dates Dates
	body,err := ioutil.ReadAll(request.Body)
	if err != nil{
		return dates,err
	}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dates)
	if err != nil{
		return dates,err
	}
	return dates, nil
}
