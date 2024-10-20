package handlers

import (
	"database/sql" // Para manipulação do banco de dados
	"encoding/json"
	"net/http"

	"github.com/sergioglesio/ott-metadata-api/pkg/models" // Certifique-se de importar o modelo
)

// Estrutura Handler
type Handler struct {
	DB *sql.DB
}

// Método para obter todos os vídeos
func (h *Handler) GetAllVideos(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT * FROM videos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var videos []models.Video
	for rows.Next() {
		var video models.Video
		if err := rows.Scan(&video.ID, &video.Title, &video.Description, &video.Cast, &video.Year, &video.Genre); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		videos = append(videos, video)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}

// Método para criar um vídeo
func (h *Handler) CreateVideo(w http.ResponseWriter, r *http.Request) {
	var video models.Video
	if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insira o vídeo no banco de dados (certifique-se de que sua tabela está configurada corretamente)
	_, err := h.DB.Exec("INSERT INTO videos (title, description, cast, year, genre) VALUES (?, ?, ?, ?, ?)",
		video.Title, video.Description, video.Cast, video.Year, video.Genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(video)
}
