package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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
	err1 := decoder.Decode(&actualResult)

	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if err1 != nil {
		t.Errorf("%v", err.Error())
	}
	if expectedCode != actualCode {
		t.Errorf("Expect code %d but get %d", expectedCode, actualCode)
	}
	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetTime_By_StartYear_1997_StartMonth_10_StartDay_16_Should_Get_Time(t *testing.T) {
	yearmonthday := YearMonthDay{
		Year:  1997,
		Month: 10,
		Day:   16,
	}
	expectedResult := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)

	actualResult := yearmonthday.GetTime()

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetTime_By_StartYear_1997_StartMonth_10_StartDay_27_Should_Get_Time(t *testing.T) {
	yearmonthday := YearMonthDay{
		Year:  1997,
		Month: 10,
		Day:   27,
	}
	expectedResult := time.Date(1997, 10, 27, 0, 0, 0, 0, time.UTC)

	actualResult := yearmonthday.GetTime()

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetTime_By_StartYear_1996_StartMonth_2_StartDay_5_Should_Get_Time(t *testing.T) {
	yearmonthday := YearMonthDay{
		Year:  1996,
		Month: 2,
		Day:   5,
	}
	expectedResult := time.Date(1996, 2, 5, 0, 0, 0, 0, time.UTC)

	actualResult := yearmonthday.GetTime()

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetTime_By_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_Time(t *testing.T) {
	yearmonthday := YearMonthDay{
		Year:  2019,
		Month: 6,
		Day:   10,
	}
	expectedResult := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := yearmonthday.GetTime()

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
