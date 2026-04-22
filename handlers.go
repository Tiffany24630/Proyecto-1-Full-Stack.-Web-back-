package main

import (
	"net/http"
	"strconv"
)

func getSeries(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if handleOptions(w, r) {
		return
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := "SELECT * FROM AnimeManga WHERE 1=1"
	args := []interface{}{}

	search := r.URL.Query().Get("search")

	if search != "" {
		query += " AND title LIKE ?"
		args = append(args, "%"+search+"%")
	}

	sort := r.URL.Query().Get("sort")
	order := r.URL.Query().Get("order")

	if sort != "" {
		if order != "desc" {
			order = "asc"
		}
		query += " ORDER BY " + sort + " " + order
	}

	query += "LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := db.Query(query, args...)

	if err != nil {
		errorResponse(w, 500, "Error fetching series")
		return
	}

	defer rows.Close()

	var list []Series

	for rows.Next() {
		var s Series
		rows.Scan(&s.ID, &s.Title, &s.Type, &s.Total, &s.Progress, &s.Image)
		list = append(list, s)
	}

	jsonResponse(w, 200, list)
}

func getSeriesByID(w http.ResponseWriter, r *http.Request, id int) {
}

func createSeries(w http.ResponseWriter, r *http.Request) {
}

func updateSeries(w http.ResponseWriter, r *http.Request) {
}

func deleteSeries(w http.ResponseWriter, r *http.Request) {
}
