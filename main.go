package main

import (
	"net/http"
)

func main() {
	initDB()

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
