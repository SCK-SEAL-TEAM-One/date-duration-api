
# date-duration-api

ข้อมูลที่เกี่ยวข้องกับ วันที่ ให้ใช้ข้อความที่สื่อไปในทางเดียวกัน โดยเริ่มจาก ปี -> เดือน -> วัน 

    ปี/เดือน/วัน
    YYYY/MM/dd
    2019/10/16
    1996/2/5

## ข้อตกลง Name Convention
1. การตั้งชื่อเทส ให้ตั้งในรูปแบบ Snake Case โดยประกอบด้วยชื่อฟังก์ชันที่ทดสอบ Input และ Output
เชื่อมชื่อฟังก์ชันกับ Input ด้วย By และเชื่อมกับ Output ด้วย Should_Return <br>
ยกตัวอย่างเช่น Test_CalculateDurationStartTimeToEndTime_By_StartYear_1997_StartMonth_10_StartDay_16_And_EndYear_2019_EndMonth_6_EndDay_10_Should_683164800_Secondes()


## ข้อตกลงใน Commit message
ADD : เพิ่ม สร้างฟังก์ชัน ไฟล์ ใหม่
DELETE : ลบ
EDIT : แก้ไข Refactor

Example : `"ADD : สร้างไฟล์ README.md และใส่ข้อตกลงในการทำงานลงไป"`

## ข้อตกลงในการเขียน Error Test Message

    if expectResult != actualResult {
    		t.Errorf("Expected %v but got %v", expectResult, actualResult)
    	}
	
## ข้อตกลงในการจัดรูปแบบ Function Unit Test
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

## วิธีการใช้งาน API

URL : https://localhost:8080/date/calculate
Method : POST

ตัวอย่าง Request Body

    {
    	"start_date" : {
    		"day":16,
    		"month":10,
    		"year": 1997
    		},
    	"end_date":{
    		"day":10,
    		"month":6,
    		"year": 2019
    	}
    }

ตัวอย่าง Response Body

    {
        "seconds": 683164800,
        "minutes": 11386080,
        "hours": 189768,
        "days": 7907,
        "start_full_date": "Thursday, 16 October 1997",
        "end_full_date": "Saturday, 10 June 0",
        "weeks": {
            "total_weeks": 1129,
            "days_of_week": 4
        },
        "months": {
            "total_months": 259,
            "days_of_month": 25
        }
    }

**Request Body**
| Key | Description |
|--|--|
| start_date | ข้อมูลวันที่สำหรับเริ่มต้น |
| end_date | ข้อมูลวันที่สำหรับวันสิ้นสุด |
| day | เลขวันที่ |
| month | เลขเดือน |
| year | เลขปี |

**Response Body** 
| Key | Description |
|--|--|
| seconds| ระยะเวลาในหน่วยวินาที |
| minutes| ระยะเวลาในหน่วยนาที |
| hours| ระยะเวลาในหน่วยชั่วโมง |
| days| ระยะเวลาในหน่วยวัน |
| start_full_date | ข้อความวันที่เริ่มต้น |
| end_full_date | ข้อความวันที่สิ้นสุด |
| weeks| ระยะเวลาหน่วยสัปดาห์ |
| total_weeks| จำนวนสัปดาห์ทั้งหมด |
| days_of_week| เศษวันที่เหลือจากสัปดาห์ |
| months| ระยะเวลาในหน่วยเดือน |
| total_months| จำนวนเดือนทั้งหมด |
| days_of_month| เศษวันที่เหลือจากเดือน |


