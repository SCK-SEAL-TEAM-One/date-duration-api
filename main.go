package main

import (
	"log"
	"net/http"
	"../date-duration-api/api"
)

func main()  {
	http.HandleFunc("/date/calculate", api.CalculateDateDuration)
	log.Fatal(http.ListenAndServe(":8080", nil))
}