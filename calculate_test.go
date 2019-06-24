package date

import(
	"testing"
)

func Test_GetFullDate_Year_1997_Month_10_Day_16_Should_Get_Thursday_16_October_1997(t *testing.T){
	date := YearMonthDay{
		year : 1997,
		month : 10,
		day : 16,
	}
	expectedResult := "Thursday, 16 October 1997"

	actualResult := date.GetFullDate()

	if expectedResult != actualResult {
		t.Errorf("Expected %s but go %s",expectedResult,actualResult)
	}
}

func Test_GetFullDate_Year_1997_Month_10_Day_27_Should_Get_Monday_27_October_1997(t *testing.T){
	date := YearMonthDay{
		year : 1997,
		month : 10,
		day : 27,
	}
	expectedResult := "Monday, 27 October 1997"

	actualResult := date.GetFullDate()

	if expectedResult != actualResult {
		t.Errorf("Expected %s but go %s",expectedResult,actualResult)
	}
}

func Test_GetFullDate_Year_1996_Month_2_Day_5_Should_Get_Monday_5_February_1996(t *testing.T){
	date := YearMonthDay{
		year : 1996,
		month : 2,
		day : 5,
	}
	expectedResult := "Monday, 5 February 1996"

	actualResult := date.GetFullDate()

	if expectedResult != actualResult {
		t.Errorf("Expected %s but go %s",expectedResult,actualResult)
	}
}

func Test_CalculateDurationStartTimeToEndTime_StartYear_1997_StartMonth_10_StartDay_16_And_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_683164800_Seconds(t *testing.T){
	start := YearMonthDay{
		year : 1997,
		month : 10,
		day : 16,
	}
	end := YearMonthDay{
		year : 2019,
		month : 6,
		day : 10,
	}
	expectedResult := Duration{
		seconds:683164800,
		minutes:11386080,
		hours:189768,
	}

	actualResult := CalculateDurationStartTimeToEndTime(start,end)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but go %v",expectedResult,actualResult)
	}
}

func Test_CalculateDurationStartTimeToEndTime_StartYear_1997_StartMonth_10_StartDay_27_And_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_682214400_Seconds(t *testing.T){
	start := YearMonthDay{
		year : 1997,
		month : 10,
		day : 27,
	}
	end := YearMonthDay{
		year : 2019,
		month : 6,
		day : 10,
	}
	expectedResult := Duration{
		seconds:682214400,
		minutes:11370240,
		hours:189504,
	}

	actualResult := CalculateDurationStartTimeToEndTime(start,end)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but go %v",expectedResult,actualResult)
	}
}

func Test_CalculateDurationStartTimeToEndTime_StartYear_1996_StartMonth_2_StartDay_5_And_EndYear_2019_EndMonth_6_EndDay_10_Should_Get_736646400_Seconds(t *testing.T){
	start := YearMonthDay{
		year : 1996,
		month : 2,
		day : 5,
	}
	end := YearMonthDay{
		year : 2019,
		month : 6,
		day : 10,
	}
	expectedResult := Duration{
		seconds:736646400,
		minutes:12277440,
		hours:204624,
	}

	actualResult := CalculateDurationStartTimeToEndTime(start,end)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but go %v",expectedResult,actualResult)
	}
}