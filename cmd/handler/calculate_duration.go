package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/SCK-SEAL-TEAM-One/date-duration-api/internal/date"
)

// RequestDate to decode
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	startTime := requestDate.StartDate.Time
	endTime := requestDate.EndDate.Time

	duration := date.CalculateDuration(startTime, endTime)
	duration.StartFullDate = date.GetFullDate(startTime)
	duration.EndFullDate = date.GetFullDate(endTime)

	err = json.NewEncoder(w).Encode(&duration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getRequestDateFromRequestBody(request *http.Request) (RequestDate, error) {
	var requestDate RequestDate
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return requestDate, err
	}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&requestDate)
	if err != nil {
		return requestDate, err
	}
	requestDate.StartDate = newYearMonthDay(requestDate.StartDate.Year, requestDate.StartDate.Month, requestDate.StartDate.Day)
	requestDate.EndDate = newYearMonthDay(requestDate.EndDate.Year, requestDate.EndDate.Month, requestDate.EndDate.Day)
	return requestDate, nil
}

func newYearMonthDay(year, month, day int) YearMonthDay {
	return YearMonthDay{
		Year:  year,
		Month: month,
		Day:   day,
		Time:  time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
	}
}
