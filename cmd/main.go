package cmd

import (
	"log"
	"net/http"

	"../internal/api"
)

func main() {
	http.HandleFunc("/date/calculate", api.CalculateDateDuration)
	log.Fatal(http.ListenAndServe(":8080", nil))
}