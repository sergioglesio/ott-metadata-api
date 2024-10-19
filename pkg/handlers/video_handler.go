package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) SearchVideos(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	genre := r.URL.Query().Get("genre")
	year := r.URL.Query().Get("year")

	var query = "SELECT * FROM videos WHERE 1=1" // Base da query
	var args []interface{}                       // Argumentos da query

	if title != "" {
		query += " AND title LIKE ?"
		args = append(args, "%"+title+"%")
	}

	if genre != "" {
		query += " AND genre = ?"
		args = append(args, genre)
	}

	if year != "" {
		query += " AND year = ?"
		args = append(args, year)
	}

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var videos []Video
	for rows.Next() {
		var video Video
		if err := rows.Scan(&video.ID, &video.Title, &video.Description, &video.Cast, &video.Year, &video.Genre); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		videos = append(videos, video)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}
