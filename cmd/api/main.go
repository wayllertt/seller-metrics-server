package api

import (
	"log"
	"net/http"

	"seller-metrics/app/app"
)

func main() {
	handler := app.NewHTTPServer()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
