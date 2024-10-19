package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const TMDbAPIKey = "9f8e4d964118e59ab7f2371dc1aef65c"

// FetchMovieDataFromTMDb faz uma requisição à API TMDb para buscar dados de filmes
func FetchMovieDataFromTMDb(w http.ResponseWriter, r *http.Request) {
	// O ID do filme pode ser passado via query param: ?movie_id=550
	movieID := r.URL.Query().Get("movie_id")
	if movieID == "" {
		http.Error(w, "Movie ID is required", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?api_key=%s", movieID, TMDbAPIKey)

	// Fazendo a requisição HTTP para a API do TMDb
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Lendo a resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Decodificando o corpo da resposta JSON
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviando o resultado como resposta ao cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
