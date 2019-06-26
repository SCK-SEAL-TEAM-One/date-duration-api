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
	dateDuration := date.CalculateDurationStartTimeToEndTime(requestDate.StartDate,requestDate.EndDate)
	duration := date.Duration{
		Seconds: dateDuration.Seconds,
		Minutes: dateDuration.Minutes,
		Hours:   dateDuration.Hours,
		Days:    dateDuration.GetDays(),
		Months:  date.CalculateMonths(requestDate.StartDate,requestDate.EndDate),
	}
	duration.Weeks = duration.GetWeeks()

	response := ResponseBody{
		StartFullDate: requestDate.StartDate.GetFullDate(),
		EndFullDate:   requestDate.EndDate.GetFullDate(),
		Duration: duration,
	}
	
	responseJSON,_ := json.Marshal(response)
	fmt.Fprintf(w, "%+v",string(responseJSON))
}

func GetRequestDateFromRequestBody(r *http.Request) RequestDate {
	body, _ := ioutil.ReadAll(r.Body)
	var requestDate RequestDate
	_ = json.Unmarshal(body, &requestDate)
	return requestDate
}
