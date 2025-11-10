package main

import (
	"log"
	"net/http"

	"seller-metrics-server/app"
)

func main() {
	handler := app.NewHTTPServer()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
