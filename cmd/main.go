package main

import (
	"log"
	"net/http"

	config "github.com/sergioglesio/ott-metadata-api/configs"
	"github.com/sergioglesio/ott-metadata-api/pkg/handlers"
)

func main() {
	// Conexão com o banco de dados
	db, err := config.DBConnection()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	// Inicializando o handler com a conexão de banco de dados
	handler := handlers.Handler{DB: db}

	// Registrando rotas
	http.HandleFunc("/videos", handler.GetAllVideos)
	http.HandleFunc("/video", handler.CreateVideo)
	// Adicione outras rotas conforme necessário

	// Iniciando o servidor
	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
