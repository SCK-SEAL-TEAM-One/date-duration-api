package date

import (
	"testing"
	"time"
)

func Test_GetFullDate_By_Year_1997_Month_10_Day_16_Should_Thursday_16_October_1997(t *testing.T) {
	expectedResult := "Thursday, 16 October 1997"
	time := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)

	actualResult := GetFullDate(time)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_1997_Month_10_Day_27_Should_Monday_27_October_1997(t *testing.T) {
	expectedResult := "Monday, 27 October 1997"
	time := time.Date(1997, 10, 27, 0, 0, 0, 0, time.UTC)

	actualResult := GetFullDate(time)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_1996_Month_2_Day_5_Should_Monday_5_February_1996(t *testing.T) {
	expectedResult := "Monday, 5 February 1996"
	time := time.Date(1996, 2, 5, 0, 0, 0, 0, time.UTC)

	actualResult := GetFullDate(time)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_2019_Month_6_Day_10_Should_Monday_10_June_2019(t *testing.T) {
	expectedResult := "Monday, 10 June 2019"
	time := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := GetFullDate(time)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetDays_By_hours_189768_Should_Days_7907(t *testing.T) {
	expectedResult := 7907
	duration := Duration{
		Hours: 189768,
	}

	duration.GetDays()
	actualResult := duration.Days

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetDays_By_hours_189504_Should_Days_7896(t *testing.T) {
	expectedResult := 7896
	duration := Duration{
		Hours: 189504,
	}

	duration.GetDays()
	actualResult := duration.Days

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetDays_By_hours_204624_Should_Days_8526(t *testing.T) {
	expectedResult := 8526
	duration := Duration{
		Hours: 204624,
	}

	duration.GetDays()
	actualResult := duration.Days

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetWeeks_By_Days_7907_Should_TotalWeeks_1129_DaysOfWeek_4(t *testing.T) {
	expectedResult := Weeks{
		TotalWeeks: 1129,
		DaysOfWeek: 4,
	}
	duration := Duration{
		Days: 7907,
	}

	duration.GetWeeks()
	actualResult := duration.Weeks

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetWeeks_By_Days_7896_Should_TotalWeeks_1128_DaysOfWeek_0(t *testing.T) {
	expectedResult := Weeks{
		TotalWeeks: 1128,
		DaysOfWeek: 0,
	}
	duration := Duration{
		Days: 7896,
	}

	duration.GetWeeks()
	actualResult := duration.Weeks

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetWeeks_By_Days_8526_Should_TotalWeeks_1218_DaysOfWeek_0(t *testing.T) {
	expectedResult := Weeks{
		TotalWeeks: 1218,
		DaysOfWeek: 0,
	}
	duration := Duration{
		Days: 8526,
	}

	duration.GetWeeks()
	actualResult := duration.Weeks

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetMonths_By_StartYear_1997_StartMonth_10_StartDay_16_Should_TotalMonths_259_DaysOfMonth_25(t *testing.T) {
	expectedResult := Months{
		TotalMonths: 259,
		DaysOfMonth: 25,
	}
	StartTime := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)
	EndTime := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := GetMonths(StartTime, EndTime)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetMonths_By_StartYear_1997_StartMonth_10_StartDay_27_Should_TotalMonths_259_DaysOfMonth_14(t *testing.T) {
	expectedResult := Months{
		TotalMonths: 259,
		DaysOfMonth: 14,
	}
	StartTime := time.Date(1997, 10, 27, 0, 0, 0, 0, time.UTC)
	EndTime := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := GetMonths(StartTime, EndTime)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetMonths_By_StartYear_1996_StartMonth_2_StartDay_5_Should_TotalMonths_280_DaysOfMonth_5(t *testing.T) {
	expectedResult := Months{
		TotalMonths: 280,
		DaysOfMonth: 5,
	}
	StartTime := time.Date(1996, 2, 5, 0, 0, 0, 0, time.UTC)
	EndTime := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := GetMonths(StartTime, EndTime)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CalculateDuration_By_StartYear_1997_StartMonth_10_StartDay_16_And_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_Duration_7907_Days(t *testing.T) {
	expectedResult := Duration{
		Days:    7907,
		Hours:   189768,
		Minutes: 11386080,
		Seconds: 683164800,
		Months: Months{
			TotalMonths: 259,
			DaysOfMonth: 25,
		},
		Weeks: Weeks{
			TotalWeeks: 1129,
			DaysOfWeek: 4,
		},
	}
	startTime := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := CalculateDuration(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CalculateDuration_By_StartYear_1997_StartMonth_10_StartDay_27_And_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_Duration_7896_Days(t *testing.T) {
	expectedResult := Duration{
		Days:    7896,
		Hours:   189504,
		Minutes: 11370240,
		Seconds: 682214400,
		Months: Months{
			TotalMonths: 259,
			DaysOfMonth: 14,
		},
		Weeks: Weeks{
			TotalWeeks: 1128,
			DaysOfWeek: 0,
		},
	}
	startTime := time.Date(1997, 10, 27, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := CalculateDuration(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CalculateDuration_By_StartYear_1996_StartMonth_2_StartDay_5_And_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_Duration_7907_Days(t *testing.T) {
	expectedResult := Duration{
		Days:    8526,
		Hours:   204624,
		Minutes: 12277440,
		Seconds: 736646400,
		Months: Months{
			TotalMonths: 280,
			DaysOfMonth: 5,
		},
		Weeks: Weeks{
			TotalWeeks: 1218,
			DaysOfWeek: 0,
		},
	}
	startTime := time.Date(1996, 2, 5, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := CalculateDuration(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
