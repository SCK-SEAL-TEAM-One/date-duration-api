package date

import (
	"testing"
)

func Test_GetFullDate_By_Year_1997_Month_10_Day_16_Should_Get_Thursday_16_October_1997(t *testing.T) {
	date := YearMonthDay{
		Year:  1997,
		Month: 10,
		Day:   16,
	}
	expectedResult := "Thursday, 16 October 1997"

	actualResult := date.GetFullDate()

	if expectedResult != actualResult {
		t.Errorf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_1997_Month_10_Day_27_Should_Get_Monday_27_October_1997(t *testing.T) {
	date := YearMonthDay{
		Year:  1997,
		Month: 10,
		Day:   27,
	}
	expectedResult := "Monday, 27 October 1997"

	actualResult := date.GetFullDate()

	if expectedResult != actualResult {
		t.Errorf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func Test_GetFullDate_By_Year_1996_Month_2_Day_5_Should_Get_Monday_5_February_1996(t *testing.T) {
	date := YearMonthDay{
		Year:  1996,
		Month: 2,
		Day:   5,
	}
	expectedResult := "Monday, 5 February 1996"

	actualResult := date.GetFullDate()

	if expectedResult != actualResult {
		t.Errorf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func Test_CalculateDurationStartTimeToEndTime_By_StartYear_1997_StartMonth_10_StartDay_16_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_683164800_Seconds(t *testing.T) {
	startDate := YearMonthDay{
		Year:  1997,
		Month: 10,
		Day:   16,
	}
	endDate := YearMonthDay{
		Year:  2019,
		Month: 6,
		Day:   10,
	}
	expectedResult := 683164800

	actualResult := CalculateDurationStartTimeToEndTime(startDate, endDate).Seconds

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_CalculateDurationStartTimeToEndTime_By_StartYear_1997_StartMonth_10_StartDay_27_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_682214400_Seconds(t *testing.T) {
	startDate := YearMonthDay{
		Year:  1997,
		Month: 10,
		Day:   27,
	}
	endDate := YearMonthDay{
		Year:  2019,
		Month: 6,
		Day:   10,
	}
	expectedResult := 682214400

	actualResult := CalculateDurationStartTimeToEndTime(startDate, endDate).Seconds

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_CalculateDurationStartTimeToEndTime_By_StartYear_1996_StartMonth_2_StartDay_5_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_736646400_Seconds(t *testing.T) {
	startDate := YearMonthDay{
		Year:  1996,
		Month: 2,
		Day:   5,
	}
	endDate := YearMonthDay{
		Year:  2019,
		Month: 6,
		Day:   10,
	}
	expectedResult := 736646400

	actualResult := CalculateDurationStartTimeToEndTime(startDate, endDate).Seconds

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetDays_By_Hours_189768_Should_Get_7907_Days(t *testing.T) {
	duration := Duration{
		Hours: 189768,
	}
	expectedResult := 7907

	actualResult := duration.GetDays()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetDays_By_Hours_189504_Should_Get_7896_Days(t *testing.T) {
	duration := Duration{
		Hours: 189504,
	}
	expectedResult := 7896

	actualResult := duration.GetDays()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetDays_By_Hours_204264_Should_Get_8526_Days(t *testing.T) {
	duration := Duration{
		Hours: 204624,
	}
	expectedResult := 8526

	actualResult := duration.GetDays()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetWeeks_By_Days_7907_Should_Get_1129_Weeks_4_Days(t *testing.T) {
	duration := Duration{
		Days: 7907,
	}
	expectedResult := WeeksDays{
		Weeks: 1129,
		Days:  4,
	}

	actualResult := duration.GetWeeks()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetWeeks_By_Days_7896_Should_Get_1128_Weeks_0_Days(t *testing.T) {
	duration := Duration{
		Days: 7896,
	}
	expectedResult := WeeksDays{
		Weeks: 1128,
		Days:  0,
	}

	actualResult := duration.GetWeeks()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetWeeks_By_Days_8526_Should_Get_1218_Weeks_0_Days(t *testing.T) {
	duration := Duration{
		Days: 8526,
	}
	expectedResult := WeeksDays{
		Weeks: 1218,
		Days:  0,
	}

	actualResult := duration.GetWeeks()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_CalculateMonths_By_StartYear_1997_StartMonth_10_StartDay_16_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_259_Months_25_Days(t *testing.T) {
	startDate := YearMonthDay{
		Year:  1997,
		Month: 10,
		Day:   16,
	}
	endDate := YearMonthDay{
		Year:  2019,
		Month: 6,
		Day:   10,
	}
	expectedResult := MonthsDays{
		Months: 259,
		Days:   25,
	}

	actualResult := CalculateMonths(startDate, endDate)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_CalculateMonths_By_StartYear_1997_StartMonth_10_StartDay_27_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_259_Months_14_Days(t *testing.T) {
	startDate := YearMonthDay{
		Year:  1997,
		Month: 10,
		Day:   27,
	}
	endDate := YearMonthDay{
		Year:  2019,
		Month: 6,
		Day:   10,
	}
	expectedResult := MonthsDays{
		Months: 259,
		Days:   14,
	}

	actualResult := CalculateMonths(startDate, endDate)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_CalculateMonths_By_StartYear_1996_StartMonth_2_StartDay_5_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_280_Months_5_Days(t *testing.T) {
	startDate := YearMonthDay{
		Year:  1996,
		Month: 2,
		Day:   5,
	}
	endDate := YearMonthDay{
		Year:  2019,
		Month: 6,
		Day:   10,
	}
	expectedResult := MonthsDays{
		Months: 280,
		Days:   5,
	}

	actualResult := CalculateMonths(startDate, endDate)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}
