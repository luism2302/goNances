package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/luism2302/goNances/database/sqlc"
	"github.com/luism2302/goNances/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error: Couldn't load .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Couldn't find PORT in .env file")
	}

	dbString := os.Getenv("DB_STRING")
	if dbString == "" {
		log.Fatal("Couldn't find DB_STRING in .env file")
	}

	conn, err := pgx.Connect(context.Background(), dbString)
	if err != nil {
		log.Fatalf("Couln't connect to db: %v", err)
	}
	defer conn.Close(context.Background())

	queries := sqlc.New(conn)

	cfg := handlers.NewConfig(queries)

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.Handle("/", handlers.MakeHandler(handlers.HandleWelcome))
	mux.Handle("GET /login", handlers.MakeHandler(handlers.HandleWelcomeLogin))
	mux.Handle("GET /signup", handlers.MakeHandler(handlers.HandleWelcomeSignUp))
	mux.Handle("POST /login", handlers.MakeHandler(cfg.HandleLogin))
	mux.Handle("POST /signup", handlers.MakeHandler(cfg.HandleUsersCreate))
	mux.Handle("GET /api/resetUsers", handlers.MakeHandler(cfg.HandleUsersDelete))

	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	//handlers
	fmt.Printf("Listening and serving on port: %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
