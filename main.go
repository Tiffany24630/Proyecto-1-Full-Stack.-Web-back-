package main

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	initDB()

	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/openapi.yaml"),
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

	http.ListenAndServe(":8080", nil)
}
