package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jackc/pgx/v5" // O driver que você acabou de instalar

	"github.com/leonardo-Trex/kuasys-backend-go/internal/db"
)

func main() {

	ctx := context.Background()
	connStr := "postgres://kuasys:kuasys@localhost:5432/kuasysdb?sslmode=disable"

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Não foi possível conectar ao banco de dados: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	err = conn.Ping(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro de ping no banco de dados: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Conexão com o PostgreSQL realizada com sucesso!")

	queries := db.New(conn)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World com conexão ao banco de dados ativa!"))
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := queries.ListUsers(r.Context())
		if err != nil {
			http.Error(w, "Erro ao buscar usuários do banco", http.StatusInternalServerError)
			return
		}

		if users == nil {
			users = []db.User{}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	})

	port := ":8000"

	err = http.ListenAndServe(port, r)
	fmt.Printf("Servidor rodando na porta %s...\n", port)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
