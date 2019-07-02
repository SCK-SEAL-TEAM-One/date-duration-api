package main

import (
	"date-duration-api/cmd/handler"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/date/calculate",handler.CalculateDuration)
	log.Fatal(http.ListenAndServe(":8080",nil))
	
}
