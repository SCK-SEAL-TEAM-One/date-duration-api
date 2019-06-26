package date

import (
	"testing"
)

func Test_GetFullDate_By_Year_1997_Month_10_Day_16_Should_Get_Thursday_16_October_1997(t *testing.T) {
	date := YearMonthDay{
		year:  1997,
		month: 10,
		day:   16,
	}
	expectedResult := "Thursday, 16 October 1997"

	actualResult := GetFullDate(date)

	if expectedResult != actualResult {
		t.Errorf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_1997_Month_10_Day_27_Should_Get_Monday_27_October_1997(t *testing.T) {
	date := YearMonthDay{
		year:  1997,
		month: 10,
		day:   27,
	}
	expectedResult := "Monday, 27 October 1997"

	actualResult := GetFullDate(date)

	if expectedResult != actualResult {
		t.Errorf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_1996_Month_2_Day_5_Should_Get_Monday_5_February_1996(t *testing.T) {
	date := YearMonthDay{
		year:  1996,
		month: 2,
		day:   5,
	}
	expectedResult := "Monday, 5 February 1996"

	actualResult := GetFullDate(date)

	if expectedResult != actualResult {
		t.Errorf("Expected %s but got %s", expectedResult, actualResult)
	}
}
