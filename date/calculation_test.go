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

	actualResult := date.GetFullDate()

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

	actualResult := date.GetFullDate()

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

	actualResult := date.GetFullDate()

	if expectedResult != actualResult {
		t.Errorf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func Test_CalculateDurationStartTimeToEndTime_By_StartYear_1997_StartMonth_10_StartDay_16_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_683164800_Seconds_11386080_Minutes_189768_Hours(t *testing.T) {
	startDate := YearMonthDay{
		year:  1997,
		month: 10,
		day:   16,
	}
	endDate := YearMonthDay{
		year:  2019,
		month: 6,
		day:   10,
	}
	expectedResult := Duration{
		seconds: 683164800,
		minutes: 11386080,
		hours:   189768,
	}

	actualResult := CalculateDurationStartTimeToEndTime(startDate, endDate)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_CalculateDurationStartTimeToEndTime_By_StartYear_1997_StartMonth_10_StartDay_27_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_682214400_Seconds_11370240_Minutes_189504_Hours(t *testing.T) {
	startDate := YearMonthDay{
		year:  1997,
		month: 10,
		day:   27,
	}
	endDate := YearMonthDay{
		year:  2019,
		month: 6,
		day:   10,
	}
	expectedResult := Duration{
		seconds: 682214400,
		minutes: 11370240,
		hours:   189504,
	}

	actualResult := CalculateDurationStartTimeToEndTime(startDate, endDate)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_CalculateDurationStartTimeToEndTime_By_StartYear_1996_StartMonth_2_StartDay_5_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_736646400_Seconds_12277440_Minutes_204624_Hours(t *testing.T) {
	startDate := YearMonthDay{
		year:  1996,
		month: 2,
		day:   5,
	}
	endDate := YearMonthDay{
		year:  2019,
		month: 6,
		day:   10,
	}
	expectedResult := Duration{
		seconds: 736646400,
		minutes: 12277440,
		hours:   204624,
	}

	actualResult := CalculateDurationStartTimeToEndTime(startDate, endDate)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetDays_By_Hours_189768_Should_Get_7907_Days(t *testing.T) {
	duration := Duration{
		hours: 189768,
	}
	expectedResult := 7907

	actualResult := duration.GetDays()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetDays_By_Hours_189504_Should_Get_7896_Days(t *testing.T) {
	duration := Duration{
		hours: 189504,
	}
	expectedResult := 7896

	actualResult := duration.GetDays()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetDays_By_Hours_204264_Should_Get_8526_Days(t *testing.T) {
	duration := Duration{
		hours: 204624,
	}
	expectedResult := 8526

	actualResult := duration.GetDays()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetWeeks_By_Days_7907_Should_Get_1129_Weeks_4_Days(t *testing.T) {
	duration := Duration{
		days: 7907,
	}
	expectedResult := WeeksDays{
		weeks: 1129,
		days: 4,
	}

	actualResult := duration.GetWeeks()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetWeeks_By_Days_7896_Should_Get_1128_Weeks_0_Days(t *testing.T) {
	duration := Duration{
		days: 7896,
	}
	expectedResult := WeeksDays{
		weeks: 1128,
		days: 0,
	}

	actualResult := duration.GetWeeks()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetWeeks_By_Days_8526_Should_Get_1218_Weeks_0_Days(t *testing.T) {
	duration := Duration{
		days: 8526,
	}
	expectedResult := WeeksDays{
		weeks: 1218,
		days: 0,
	}

	actualResult := duration.GetWeeks()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}