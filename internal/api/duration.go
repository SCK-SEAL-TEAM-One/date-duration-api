package api

import (
	"../date"
	"bytes"
	"encoding/json"
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
	response := ResponseBody{
		StartFullDate: requestDate.StartDate.GetFullDate(),
		EndFullDate:   requestDate.EndDate.GetFullDate(),
		Duration:      date.CalculateDurationStartTimeToEndTime(requestDate.StartDate, requestDate.EndDate),
	}

	enc := json.NewEncoder(w)
	enc.Encode(response)
}

func GetRequestDateFromRequestBody(r *http.Request) RequestDate {
	body, _ := ioutil.ReadAll(r.Body)
	var requestDate RequestDate
	decode := json.NewDecoder(bytes.NewReader(body))
	decode.Decode(&requestDate)
	return requestDate
}
