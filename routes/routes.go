package routes

import (
	"net/http"

	"github.com/api_loja/handlers"
)

func HandleRequest() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/novo", handlers.Novo)
	http.HandleFunc("/inserir", handlers.Inserir)
	http.HandleFunc("/deletar", handlers.Deletar)
	http.HandleFunc("/editar", handlers.Editar)
	http.HandleFunc("/atualizar", handlers.Atualizar)
}
