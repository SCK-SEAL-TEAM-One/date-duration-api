package date

import (
	"testing"
	"time"
)

func Test_GetDays_By_189768_Hour_Should_Get_7907_Day(t *testing.T) {
	expectedResult := 7907

	actualResult := getDays(189768)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetDays_By_189504_Hour_Should_Get_7896_Day(t *testing.T) {
	expectedResult := 7896

	actualResult := getDays(189504)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetDays_By_204624_Hour_Should_Get_8526_Day(t *testing.T) {
	expectedResult := 8526

	actualResult := getDays(204624)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetWeek_By_7907_Day_Should_Get_1129_TotalWeeks_And_4_DayOfWeek(t *testing.T) {
	expectedResult := Weeks{TotalWeeks: 1129, DaysOfWeek: 4}

	actualResult := getWeeks(7907)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetWeek_By_7896_Day_Should_Get_1128_TotalWeeks_And_0_DayOfWeek(t *testing.T) {
	expectedResult := Weeks{TotalWeeks: 1128, DaysOfWeek: 0}

	actualResult := getWeeks(7896)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetWeek_By_8526_Day_Should_Get_1218_TotalWeeks_And_0_DayOfWeek(t *testing.T) {
	expectedResult := Weeks{TotalWeeks: 1218, DaysOfWeek: 0}

	actualResult := getWeeks(8526)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetMonths_By_StartTime_Year_1997_Month_10_Day_16_And_EndTime_Year_2019_Month_6_Day_10_Should_Get_259_Months_And_25_Days(t *testing.T) {
	expectedResult := Months{TotalMonths: 259, DaysOfMonth: 25}
	startTime := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := getMonths(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetMonths_By_StartTime_Year_1997_Month_10_Day_27_And_EndTime_Year_2019_Month_6_Day_10_Should_Get_259_Months_And_14_Days(t *testing.T) {
	expectedResult := Months{TotalMonths: 259, DaysOfMonth: 14}
	startTime := time.Date(1997, 10, 27, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := getMonths(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetMonths_By_StartTime_Year_1996_Month_2_Day_5_And_EndTime_Year_2019_Month_6_Day_10_Should_Get_280_Months_And_5_Days(t *testing.T) {
	expectedResult := Months{TotalMonths: 280, DaysOfMonth: 5}
	startTime := time.Date(1996, 2, 5, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := getMonths(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetMonths_By_StartTime_Year_2019_Month_5_Day_10_And_EndTime_Year_2019_Month_6_Day_10_Should_Get_1_Months_And_0_Days(t *testing.T) {
	expectedResult := Months{TotalMonths: 1, DaysOfMonth: 0}
	startTime := time.Date(2019, 5, 10, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := getMonths(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetMonths_By_StartTime_Year_2019_Month_2_Day_10_And_EndTime_Year_2019_Month_3_Day_10_Should_Get_1_Months_And_0_Days(t *testing.T) {
	expectedResult := Months{TotalMonths: 1, DaysOfMonth: 0}
	startTime := time.Date(2019, 2, 10, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 3, 10, 0, 0, 0, 0, time.UTC)

	actualResult := getMonths(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_1997_Month_10_Day_16_Should_Get_Thursday_16_October_1997(t *testing.T) {
	date := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)
	expectedResult := "Thursday, 16 October 1997"

	actualResult := getFullDate(date)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_1997_Month_10_Day_27_Should_Get_Monday_27_October_1997(t *testing.T) {
	date := time.Date(1997, 10, 27, 0, 0, 0, 0, time.UTC)
	expectedResult := "Monday, 27 October 1997"

	actualResult := getFullDate(date)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_1996_Month_2_Day_5_Should_Get_Monday_5_February_1996(t *testing.T) {
	date := time.Date(1996, 2, 5, 0, 0, 0, 0, time.UTC)
	expectedResult := "Monday, 5 February 1996"

	actualResult := getFullDate(date)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_2019_Month_6_Day_10_Should_Get_Monday_10_June_2019(t *testing.T) {
	date := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)
	expectedResult := "Monday, 10 June 2019"

	actualResult := getFullDate(date)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}
