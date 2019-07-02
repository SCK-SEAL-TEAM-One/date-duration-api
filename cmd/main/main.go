package main

import (
	"date-duration-api/cmd/handler"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const defaultHttpPort = 8080

func main() {
	httpPrt := flag.Int("http", defaultHttpPort, "")
	flag.Parse()
	addr := fmt.Sprintf(":%d", *httpPrt)

	http.HandleFunc("/date/calculate", handler.CalculateDuration)
	log.Fatal(http.ListenAndServe(addr, nil))
}
