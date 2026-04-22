package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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
	enableCORS(w)

	if handleOptions(w, r) {
		return
	}

	idstr := strings.TrimPrefix(r.URL.Path, "/series/")
	id, _ = strconv.Atoi(idstr)

	var s Series
	err := db.QueryRow("SELECT * FROM AnimeManga WHERE id = ?", id).Scan(&s.ID, &s.Title, &s.Type, &s.Total, &s.Progress, &s.Image)

	if err != nil {
		errorResponse(w, 404, "Series not found")
		return
	}

	jsonResponse(w, 200, s)
}

func createSeries(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if handleOptions(w, r) {
		return
	}

	var s Series
	json.NewDecoder(r.Body).Decode(&s)

	if s.Title == "" {
		errorResponse(w, 400, "title is required")
		return
	}

	res, err := db.Exec(
		"INSERT INTO AnimeManga(title, type, total_episodes, watched_episodes, image) VALUES (?, ?, ?, ?, ?)",
		s.Title, s.Type, s.TotalEpisodes, s.WatchedEpisodes, s.Image,
	)

	if err != nil {
		errorResponse(w, 500, "insert error")
		return
	}

	id, _ := res.LastInsertId()
	s.ID = int(id)

	jsonResponse(w, 201, s)
}

func updateSeries(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if handleOptions(w, r) {
		return
	}

	idstr := strings.TrimPrefix(r.URL.Path, "/series/")
	id, _ = strconv.Atoi(idstr)

	var s Series
	json.NewDecoder(r.Body).Decode(&s)

	_, err := db.Exec(
		"UPDATE AnimeManga SET title = ?, type = ?, total_episodes = ?, watched_episodes = ?, image = ? WHERE id = ?",
		s.Title, s.Type, s.TotalEpisodes, s.WatchedEpisodes, s.Image, id,
	)

	if err != nil {
		errorResponse(w, 500, "update error")
		return
	}

	jsonResponse(w, 200, s)
}

func deleteSeries(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if handleOptions(w, r) {
		return
	}

	idstr := strings.TrimPrefix(r.URL.Path, "/series/")
	id, _ = strconv.Atoi(idstr)

	res, _ := db.Exec("DELETE FROM AnimeManga WHERE id = ?", id)
	rows, _ := res.RowsAffected()

	if rows == 0 {
		errorResponse(w, 404, "Series not found")
		return
	}

	w.WriteHeader(204)

}
