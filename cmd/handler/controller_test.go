package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SCK-SEAL-TEAM-One/date-duration-api/internal/date"
)

func Test_CalculateDuration_Between_Mo_Birthday_And_10_June_2019(t *testing.T) {
	requestDate := `{"start_date":{"year":1997,"month":10,"day":16},"end_date":{"year":2019,"month":6,"day":10}}`
	request := httptest.NewRequest("POST", "http://localhost:8080/date/calculate", strings.NewReader(requestDate))
	recorder := httptest.NewRecorder()
	expectedResult := date.Duration{
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
	expectedCode := 200

	CalculateDuration(recorder, request)
	response := recorder.Result()
	body, err := ioutil.ReadAll(response.Body)
	actualCode := response.StatusCode
	var actualResult date.Duration
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.Decode(&actualResult)

	if err != nil {
		t.Errorf("%v", err)
	}
	if expectedCode != actualCode {
		t.Errorf("Expect code %d but get %d", expectedCode, actualCode)
	}
	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}