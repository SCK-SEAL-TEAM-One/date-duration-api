package date

import (
	"testing"
	"time"
)

func Test_GetWeek_By_7907_Day_Should_Get_1129_TotalWeeks_And_4_DayOfWeek(t *testing.T) {
	expectedResult := Weeks{1129, 4}

	actualResult := GetWeeks(7907)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetWeek_By_7896_Day_Should_Get_1128_TotalWeeks_And_0_DayOfWeek(t *testing.T) {
	expectedResult := Weeks{1128, 0}

	actualResult := GetWeeks(7896)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetWeek_By_8526_Day_Should_Get_1218_TotalWeeks_And_0_DayOfWeek(t *testing.T) {
	expectedResult := Weeks{1218, 0}

	actualResult := GetWeeks(8526)

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
