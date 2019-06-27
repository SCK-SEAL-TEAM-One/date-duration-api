package date

import (
	"testing"
	"time"
)

func Test_GetFullDate_By_Year_1997_Month_10_Day_16_Should_Thursday_16_October_1997(t *testing.T) {
	expectedResult := "Thursday, 16 October 1997"
	time := time.Date(1997,10,16,0,0,0,0,time.UTC)

	actualResult := GetFullDate(time)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_1997_Month_10_Day_27_Should_Monday_27_October_1997(t *testing.T) {
	expectedResult := "Monday, 27 October 1997"
	time := time.Date(1997,10,27,0,0,0,0,time.UTC)

	actualResult := GetFullDate(time)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_1996_Month_2_Day_5_Should_Monday_5_February_1996(t *testing.T) {
	expectedResult := "Monday, 5 February 1996"
	time := time.Date(1996,2,5,0,0,0,0,time.UTC)

	actualResult := GetFullDate(time)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_2019_Month_6_Day_10_Should_Monday_10_June_2019(t *testing.T) {
	expectedResult := "Monday, 10 June 2019"
	time := time.Date(2019,6,10,0,0,0,0,time.UTC)

	actualResult := GetFullDate(time)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
