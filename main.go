package main

import (
	"log"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	initDB()

	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/openapi.yaml")
	))

	http.HandleFunc("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./openapi.yaml")
	})

	http.HandleFunc("/series", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getSeries(w, r)

		case "POST":
			createSeries(w, r)

		case "PUT":
			updateSeries(w, r)

		case "DELETE":
			deleteSeries(w, r)

		default:
			http.NotFound(w, r)

		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Running on port:", port)
	http.ListenAndServe(":"+port, nil)
}
