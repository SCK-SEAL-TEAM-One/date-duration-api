package main

import (
	"log"
	"net/http"

	"github.com/SCK-SEAL-TEAM-One/date-duration-api/cmd/handler"
)

func main() {
	http.HandleFunc("/date/calculate", handler.CalculateDuration)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
