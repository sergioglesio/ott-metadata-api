package handlers

import (
	"database/sql"
	"net/http"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) GetAllVideos(w http.ResponseWriter, r *http.Request) {
	// SQL para buscar todos os vídeos
}

func (h *Handler) GetVideoByID(w http.ResponseWriter, r *http.Request) {
	// SQL para buscar vídeo por ID
}

func (h *Handler) CreateVideo(w http.ResponseWriter, r *http.Request) {
	// SQL para inserir novo vídeo
}

func (h *Handler) UpdateVideo(w http.ResponseWriter, r *http.Request) {
	// SQL para atualizar vídeo
}

func (h *Handler) DeleteVideo(w http.ResponseWriter, r *http.Request) {
	// SQL para deletar vídeo
}
