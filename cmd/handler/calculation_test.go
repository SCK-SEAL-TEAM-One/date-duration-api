package handler

import (
	"bytes"
	"date-duration-api/internal/date"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func CalculateDuration_Between_Mo_Birthday_And_10_June_2019(t *testing.T)  {
	expectedResult := date.Duration{
		Seconds:683164800,
		Minutes :11386080,
		Hours:189768,
		Days :7907,
		Weeks :date.Weeks{
			TotalWeeks :1129,
			DaysOfWeek:4,
		},
		Months :date.Months{
			TotalMonths:259,
			DaysOfMonth:25,
		},
		StartFullDate :"Thursday, 16 October 1997",
		EndFullDate :"Monday, 10 June 2019",
	}
	w := httptest.NewRecorder()
	request := httptest.NewRequest("POST","http://localhost:8080/date/calculate",strings.NewReader(`{"start_date":{"year":1997,"month":10,"day":16},"end_date":{"year":2019,"month":6,"day":10}}`))
	CalculateDuration(w,request)
	response := w.Result()
	if actualHttpStatus := response.StatusCode; http.StatusOK != actualHttpStatus{
		t.Fatalf("Exception HTTP Status %d",actualHttpStatus)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		t.Fatalf("Exeption %v",err.Error())
	}

	var actualResult date.Duration
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&actualResult)
	if err != nil{
		t.Fatalf("Exeption %v",err.Error())
	}

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_DecodeDate_By_StartYear_1997_StartMonth_10_StartDay_16_And_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_Dates(t *testing.T)  {
	request := httptest.NewRequest("POST","http://localhost:8080/date/calculate",strings.NewReader(`{"start_date":{"year":1997,"month":10,"day":16},"end_date":{"year":2019,"month":6,"day":10}}`))
	expectedResult := Dates{
		StartDate:YearMonthDay{
			Year:1997,
			Month:10,
			Day:16,
		},
		EndDate:YearMonthDay{
			Year:2019,
			Month:6,
			Day:10,
		},
	}

	actualResult,err := decodeDate(request);
	if err != nil{
		t.Fatalf("Exeption %v",err.Error())
	}

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_DecodeDate_By_StartYear_1997_StartMonth_10_StartDay_27_And_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_Dates(t *testing.T)  {
	request := httptest.NewRequest("POST","http://localhost:8080/date/calculate",strings.NewReader(`{"start_date":{"year":1997,"month":10,"day":27},"end_date":{"year":2019,"month":6,"day":10}}`))
	expectedResult := Dates{
		StartDate:YearMonthDay{
			Year:1997,
			Month:10,
			Day:27,
		},
		EndDate:YearMonthDay{
			Year:2019,
			Month:6,
			Day:10,
		},
	}

	actualResult,err := decodeDate(request);
	if err != nil{
		t.Fatalf("Exeption %v",err.Error())
	}

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_DecodeDate_By_StartYear_1996_StartMonth_2_StartDay_5_And_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_Dates(t *testing.T)  {
	request := httptest.NewRequest("POST","http://localhost:8080/date/calculate",strings.NewReader(`{"start_date":{"year":1996,"month":2,"day":5},"end_date":{"year":2019,"month":6,"day":10}}`))
	expectedResult := Dates{
		StartDate:YearMonthDay{
			Year:1996,
			Month:2,
			Day:5,
		},
		EndDate:YearMonthDay{
			Year:2019,
			Month:6,
			Day:10,
		},
	}

	actualResult,err := decodeDate(request);
	if err != nil{
		t.Fatalf("Exeption %v",err.Error())
	}

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}