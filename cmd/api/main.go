package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World com Go-Chi!"))
	})

	port := ":8080"
	fmt.Printf("Servidor rodando com sucesso na porta %s...\n", port)

	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
