package handler

import (
	"bytes"
	"date-duration-api/internal/date"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	nanoSecStartOfDay = 0
	secStartOfDay     = 0
	minStartOfDay     = 0
	hoursStartOfDay   = 0
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
	dates, err := decodeDate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	startTime := dates.StartDate.getTime()
	endTime := dates.EndDate.getTime()
	duration := date.DurationBetween(startTime, endTime)
	err = json.NewEncoder(w).Encode(&duration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func decodeDate(request *http.Request) (Dates, error) {
	var dates Dates
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return dates, err
	}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dates)
	if err != nil {
		return dates, err
	}
	return dates, nil
}

func (yearMonthDay YearMonthDay) getTime() time.Time {
	return time.Date(yearMonthDay.Year, time.Month(yearMonthDay.Month), yearMonthDay.Day, hoursStartOfDay, minStartOfDay, secStartOfDay, nanoSecStartOfDay, time.UTC)
}
