package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/leonardo-Trex/kuasys-backend-go/internal/user"

	"github.com/jackc/pgx/v5"

	"github.com/leonardo-Trex/kuasys-backend-go/internal/db"
)

//TODO: Graceful shutdown
//TODO: Configs management with godotenv

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

	userService := user.NewService(queries)
	userHandler := user.NewHandler(userService)

	r := newChi()

	userHandler.RegisterRoutes(r)

	port := ":8000"

	fmt.Printf("Servidor rodando na porta %s...\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}

func newChi() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Kuasys rodando com arquitetura modular!"))
	})

	return r
}
