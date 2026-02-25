package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/luism2302/goNances/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error: Couldn't load .env file")
	}
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.Handle("/", handlers.MakeHandler(handlers.HandleWelcome))
	mux.Handle("GET /signup", handlers.MakeHandler(handlers.HandleSignUp))
	mux.Handle("POST /users", handlers.MakeHandler(handlers.HandleUsersCreate))
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
