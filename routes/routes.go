package routes

import (
	"net/http"

	"github.com/api_loja/handlers"
)

func HandleRequest() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/new", handlers.New)
	http.HandleFunc("/insert", handlers.Insert)
	http.HandleFunc("/delete", handlers.Delete)
	http.HandleFunc("/edit", handlers.Edit)
	http.HandleFunc("/update", handlers.Update)
}
