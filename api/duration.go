package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CalculateDateDuration(w http.ResponseWriter, r *http.Request) {

GetRequestDateFromRequestBody(r)
	fmt.Fprintf(w,`{
"start_full_date":"Thursday, 16 October 1997",
"end_full_date":"Monday, 10 June 2019",
"duration":{
"seconds":683164800,
"minutes":11386080,
"hours":189768,
"days":7907,
"weeks":{"weeks":1129,"days":4},
"months":{"months":259,"days":25}
}
}`)
}

func GetRequestDateFromRequestBody(r *http.Request){
	body,_:=ioutil.ReadAll(r.Body)
	var requestDate RequestDate
	_ = json.Unmarshal(body,&requestDate)
	fmt.Printf("%+v",requestDate)
}

type RequestDate struct {
	StartDate YearMonthDay `json:"start_date"`
	EndDate YearMonthDay `json:"end_date"`
}

type YearMonthDay struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type ResponseBody struct {
	StartFullDate string `json:"start_full_date"`
	EndFullDate string `json:"end_full_date"`
	Duration Duration `json:"duration"`
}
type Duration struct {
	Seconds int `json:"seconds"`
	Minutes int `json:"minutes"`
	Hours int `json:"hours"`
	Days int `json:"days"`
	Weeks Weeks `json:"weeks"`
	Months Months `json:"months"`
}
type Weeks struct {
	Weeks int `json:"weeks"`
	Days int `json:"days"`
}

type Months struct {
	Months int `json:"months"`
	Days int `json:"days"`
}