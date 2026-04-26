package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func seriesHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if handleOptions(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		getSeries(w, r)
	case "POST":
		createSeries(w, r)
	default:
		http.NotFound(w, r)
	}
}

func seriesByIDHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if handleOptions(w, r) {
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/series/")
	id, _ := strconv.Atoi(idStr)

	switch r.Method {
	case "GET":
		getSeriesByID(w, id)
	case "PUT":
		updateSeries(w, r, id)
	case "DELETE":
		deleteSeries(w, id)
	default:
		http.NotFound(w, r)
	}
}

func getSeries(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("q")

	query := "SELECT id, title, type, total, progress, image FROM AnimeManga WHERE 1=1"
	args := []interface{}{}

	if search != "" {
		query += " AND title LIKE ?"
		args = append(args, "%"+search+"%")
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		errorResponse(w, 500, err.Error())
		return
	}
	defer rows.Close()

	var list []Series

	for rows.Next() {
		var s Series
		if err := rows.Scan(&s.ID, &s.Title, &s.Type, &s.Total, &s.Progress, &s.Image); err != nil {
			errorResponse(w, 500, err.Error())
			return
		}
		list = append(list, s)
	}

	jsonResponse(w, 200, list)
}

func getSeriesByID(w http.ResponseWriter, id int) {
	var s Series
	err := db.QueryRow("SELECT id, title, type, total, progress, image FROM AnimeManga WHERE id = ?", id).
		Scan(&s.ID, &s.Title, &s.Type, &s.Total, &s.Progress, &s.Image)

	if err != nil {
		errorResponse(w, 404, "Not found")
		return
	}

	jsonResponse(w, 200, s)
}

func createSeries(w http.ResponseWriter, r *http.Request) {
	var s Series
	json.NewDecoder(r.Body).Decode(&s)

	res, err := db.Exec(
		"INSERT INTO AnimeManga(title, type, total, progress, image) VALUES (?, ?, ?, ?, ?)",
		s.Title, s.Type, s.Total, s.Progress, s.Image,
	)

	if err != nil {
		errorResponse(w, 500, err.Error())
		return
	}

	id, _ := res.LastInsertId()
	s.ID = int(id)

	jsonResponse(w, 201, s)
}

func updateSeries(w http.ResponseWriter, r *http.Request, id int) {
	var s Series
	json.NewDecoder(r.Body).Decode(&s)

	_, err := db.Exec(
		"UPDATE AnimeManga SET title=?, type=?, total=?, progress=?, image=? WHERE id=?",
		s.Title, s.Type, s.Total, s.Progress, s.Image, id,
	)

	if err != nil {
		errorResponse(w, 500, err.Error())
		return
	}

	jsonResponse(w, 200, s)
}

func deleteSeries(w http.ResponseWriter, id int) {
	res, _ := db.Exec("DELETE FROM AnimeManga WHERE id=?", id)
	rows, _ := res.RowsAffected()

	if rows == 0 {
		errorResponse(w, 404, "Not found")
		return
	}

	w.WriteHeader(204)
}

func updateProgress(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if handleOptions(w, r) {
		return
	}

	var data struct {
		ID       int `json:"id"`
		Progress int `json:"progress"`
	}

	json.NewDecoder(r.Body).Decode(&data)

	_, err := db.Exec("UPDATE AnimeManga SET progress=? WHERE id=?", data.Progress, data.ID)
	if err != nil {
		errorResponse(w, 500, err.Error())
		return
	}

	jsonResponse(w, 200, map[string]string{"message": "updated"})
}
