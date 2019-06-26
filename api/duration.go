package api

import (
	"../date"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RequestDate struct {
	StartDate date.YearMonthDay `json:"start_date"`
	EndDate   date.YearMonthDay `json:"end_date"`
}

type ResponseBody struct {
	StartFullDate string        `json:"start_full_date"`
	EndFullDate   string        `json:"end_full_date"`
	Duration      date.Duration `json:"duration"`
}

func CalculateDateDuration(w http.ResponseWriter, r *http.Request) {
	requestDate := GetRequestDateFromRequestBody(r)
	duration := ResponseBody{
		StartFullDate: requestDate.StartDate.GetFullDate(),
		EndFullDate:   requestDate.EndDate.GetFullDate(),
		Duration: date.Duration{
			Seconds: date.CalculateDurationStartTimeToEndTime(requestDate.StartDate,requestDate.EndDate).Seconds,
			Minutes: date.CalculateDurationStartTimeToEndTime(requestDate.StartDate,requestDate.EndDate).Minutes,
			Hours:   date.CalculateDurationStartTimeToEndTime(requestDate.StartDate,requestDate.EndDate).Hours,
			Days:    date.CalculateDurationStartTimeToEndTime(requestDate.StartDate,requestDate.EndDate).GetDays(),
			Months:  date.CalculateMonths(requestDate.StartDate,requestDate.EndDate),
		},
	}

	durationJSON,_ := json.Marshal(duration)
	fmt.Fprintf(w, "%+v",string(durationJSON))
}

func GetRequestDateFromRequestBody(r *http.Request) RequestDate {
	body, _ := ioutil.ReadAll(r.Body)
	var requestDate RequestDate
	_ = json.Unmarshal(body, &requestDate)
	return requestDate
}
