package api

import (
	"fmt"
	"net/http"
)

func CalculateDateDuration(w http.ResponseWriter, r *http.Request) {
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