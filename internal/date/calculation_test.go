package date

import "testing"

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
