# date-duration-api

ข้อมูลที่เกี่ยวข้องกับ วันที่ ให้ใช้ข้อความที่สื่อไปในทางเดียวกัน โดยเริ่มจาก ปี -> เดือน -> วัน 

ปี/เดือน/วัน<br>
YYYY/MM/dd<br>
2019/10/16<br>
1996/2/5
<br><br>
ข้อตกลง Name Convention
1. การตั้งชื่อเทส ให้ตั้งในรูปแบบ Snake Case โดยประกอบด้วยชื่อฟังก์ชันที่ทดสอบ Input และ Output
เชื่อมชื่อฟังก์ชันกับ Input ด้วย By และเชื่อมกับ Output ด้วย Should_Return <br>
ยกตัวอย่างเช่น Test_CalculateDurationStartTimeToEndTime_By_StartYear_1997_StartMonth_10_StartDay_16_And_EndYear_2019_EndMonth_6_EndDay_10_Should_683164800_Secondes()

<br><br>
ข้อตกลงใน Commit message<br>
ADD : เพิ่ม สร้างฟังก์ชัน ไฟล์ ใหม่<br>
DELETE : ลบ<br>
EDIT : แก้ไข Refactor<br>
Example "ADD : สร้างไฟล์ README.md และใส่ข้อตกลงในการทำงานลงไป"

<br><br>
ข้อตกลงในการเขียน Error Test Message

if expectResult != actualResult {<br>
		t.Errorf("Expected %v but got %v", expectResult, actualResult)<br>
	}
	

ข้อตกลงในการจัดรูปแบบ Function Unit Test
ให้ทำการ เว้นบรรทัด แยกการทำงานออกเป็น 3 ส่วนคือ 1.arrange 2.action 3.assert ตัวอย่าง

    func Test_ConventYearMonthDayToTime_By_Year_1997_Month_10_Day_16_Should_Return_Time_Date_Year_1997_Month_10_Day_16(t *testing.T) {
	    expectResult := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)
	    yearMonthDay := model.YearMonthDay{
		Year:  1997,
		Month: 10,
		Day:   16,
	    }
        /---- เว้นบรรทัดระหว่าง Arrange กับ Action ----/
	    actualResult := duration.ConventYearMonthDayToTime(yearMonthDay)
        /---- เว้นบรรทัดระหว่าง Action กับ Assert ----/
	    if expectResult != actualResult {
		    t.Errorf("Expected %v but got %v", expectResult, actualResult)
	    }
    }

